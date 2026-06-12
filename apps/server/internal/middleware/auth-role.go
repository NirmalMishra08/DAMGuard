package middlewares

import (
	"context"
	"main/internal/repository/postgres/sqlc"
	"main/internal/utils"
)

// CheckUserRole - Core reusable authorization function
func HasRole(store *sqlc.Store, ctx context.Context, requiredRole sqlc.UserRole) error {
	payload, err := GetFirebasePayloadFromContext(ctx)
	if err != nil {
		return err
	}

	user, err := store.Queries.GetUser(ctx, payload.UserId)
	if err != nil {
		return err
	}

	if user.Role != requiredRole {
		return utils.ErrUnauthorized
	}

	return nil
}
