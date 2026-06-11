package middlewares

import (
	"context"
)

// CheckUserRole - Core reusable authorization function
func HasRole(store db.Store, ctx context.Context, requiredRole db.UserRole) error {
	payload, err := GetFirebasePayloadFromContext(ctx)
	if err != nil {
		return err
	}

	role, err := store.GetUserRole(ctx, payload.UserId)
	if err != nil {
		return err
	}

	if role != requiredRole {
		return util.ErrUnauthorized
	}

	return nil
}