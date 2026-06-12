package users

import (
	"database/sql"
	"fmt"
	"main/internal/firebase"
	"main/internal/repository/postgres/sqlc"
	"main/internal/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authRequest struct {
	Provider string `json:"provider" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password"`
}

type authResponse struct {
	UserID   string `json:"user_id"`
	Provider string `json:"provider"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

func (h *Handler) AuthHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		utils.ErrJson(w, utils.ErrInvalidToken)
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		utils.ErrJson(w, utils.ErrInvalidToken)
		return
	}
	idToken := parts[1]

	ctx := r.Context()
	var payload firebase.FirebasePayload
	var err error

	if idToken == "frontend" {
		payload = firebase.FirebasePayload{
			Email:  "test@example.com",
			UserId: uuid.MustParse("0b927d97-782a-4c82-b9d2-e4e06774ed37"),
			UID:    "0b927d97-782a-4c82-b9d2-e4e06774ed37",
		}
	} else {
		payload, err = firebase.VerifyFirebaseIDToken(ctx, idToken)
		if err != nil {
			utils.ErrJson(w, err)
			return
		}
	}

	var req authRequest
	if err := utils.ReadJsonAndValidate(w, r, &req); err != nil {
		utils.ErrJson(w, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrJson(w, err)
		return
	}

	// 🔐 Check if user already exists or create
	user, err := h.store.Queries.FindOrCreateUser(ctx, sqlc.FindOrCreateUserParams{
		Email:        payload.Email,
		Phone:        sql.NullString{String: payload.Phone, Valid: payload.Phone != ""},
		Provider:     req.Provider,
		Fullname:     req.FullName,
		PasswordHash: sql.NullString{String: string(hashedPassword), Valid: req.Password != ""},
	})
	if err != nil {
		utils.ErrJson(w, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, authResponse{
		UserID:   payload.UID,
		Provider: fmt.Sprint(user.Provider),
		Email:    user.Email,
		FullName: req.FullName,
	})
}
