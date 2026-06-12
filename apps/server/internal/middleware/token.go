package middlewares

import (
	"context"
	"database/sql"
	"errors"
	"main/internal/firebase"
	"main/internal/repository/postgres/sqlc"
	"main/internal/utils"

	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus" // Ensure you have a logging package like logrus or use your logger
)

type tokenPayloadKeyType string

const TokenPayloadKey tokenPayloadKeyType = "auth-payload"

// Accepts sqlc.Queries so we can fetch internal user UUID
func TokenMiddleware(store *sqlc.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.ErrJson(w, utils.ErrInvalidToken)
				logrus.Error("Authorization header missing")
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				utils.ErrJson(w, utils.ErrInvalidToken)
				logrus.Error("Invalid Authorization header format")
				return
			}

			idToken := parts[1]

			// For testing purposes, we bypass Firebase token verification (dev mode)
			if idToken == "frontend" {
				ctx := context.WithValue(r.Context(), TokenPayloadKey, firebase.FirebasePayload{
					Email:  "test@example.com",
					UserId: uuid.MustParse("0b927d97-782a-4c82-b9d2-e4e06774ed37"),
					UID:    "0b927d97-782a-4c82-b9d2-e4e06774ed37",
				})
				logrus.Info("Frontend bypass: using test user payload")
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			// Verify Firebase ID token
			fbPayload, err := firebase.VerifyFirebaseIDToken(r.Context(), idToken)
			if err != nil {
				utils.ErrJson(w, utils.ErrInvalidToken)
				logrus.WithError(err).Error("Failed to verify Firebase ID token")
				return
			}

			logrus.Infof("Firebase token verified. Email: %s, UserID: %s", fbPayload.Email, fbPayload.UserId)

			user, err := store.Queries.FindOrCreateUser(r.Context(), sqlc.FindOrCreateUserParams{
				Email:        fbPayload.Email,
				Fullname:     fbPayload.Fullname,                       // Use the correct type for Fullname
				PasswordHash: sql.NullString{String: "", Valid: false}, // Password is not used here
				Phone:        sql.NullString{String: fbPayload.Phone, Valid: fbPayload.Phone != ""},
				Provider:     fbPayload.Provider, // or extract from token
			})
			if err != nil {
				utils.ErrJson(w, err)
				logrus.WithError(err).Error("Failed to find or create user")
				return
			}

			logrus.Infof("User found/created. User ID: %v", user.ID)

			// Update the Firebase payload with the internal user ID
			fbPayload.UserId = user.ID

			logrus.Debugf("User ID set in payload: %s", fbPayload.UserId)

			// Set the payload in the context
			ctx := context.WithValue(r.Context(), TokenPayloadKey, fbPayload)
			logrus.Info("Payload set to context")

			// Proceed with the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetFirebasePayloadFromContext(ctx context.Context) (firebase.FirebasePayload, error) {
	raw := ctx.Value(TokenPayloadKey)
	if raw == nil {
		return firebase.FirebasePayload{}, errors.New("missing auth token payload in context")
	}

	payload, ok := raw.(firebase.FirebasePayload)
	if !ok {
		return firebase.FirebasePayload{}, errors.New("invalid auth token payload type")
	}

	return payload, nil
}
