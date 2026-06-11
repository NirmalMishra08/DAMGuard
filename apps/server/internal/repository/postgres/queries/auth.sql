-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: FindOrCreateUser :one
INSERT INTO users (email, provider, phone, fullname, password_hash, created_at)
VALUES ($1, $2, $3, $4, $5, NOW())
ON CONFLICT (email)
DO UPDATE SET
    phone = EXCLUDED.phone,
    provider = EXCLUDED.provider,
    updated_at = NOW()
RETURNING id, email, fullname, password_hash, phone, role, provider;
