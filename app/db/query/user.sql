-- name: CreateUser :one
INSERT INTO users (email, name, password_hash)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET email = $2,
    name = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :one
UPDATE users
SET deleted_at = NOW()
WHERE id = $1
RETURNING *;
