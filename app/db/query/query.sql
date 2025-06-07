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

-- name: CreateExpense :one
INSERT INTO expenses (user_id, sub_category_id, expense_date, amount, description)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetExpense :one
SELECT * FROM expenses WHERE id = $1 LIMIT 1;

-- name: GetAllExpenses :many
SELECT * FROM expenses WHERE user_id = $1 ORDER BY expense_date DESC;

-- name: UpdateExpense :one
UPDATE expenses
SET sub_category_id = $2,
	expense_date = $3,
	amount = $4,
	description = $5,
	updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteExpense :one
UPDATE expenses
SET deleted_at = NOW()
WHERE id = $1
RETURNING *;

-- name: CreateIncome :one
INSERT INTO incomes (user_id, sub_category_id, income_date, amount, description)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetIncome :one
SELECT * FROM incomes WHERE id = $1 LIMIT 1;

-- name: GetAllIncomes :many
SELECT * FROM incomes WHERE user_id = $1 ORDER BY income_date DESC;

-- name: UpdateIncome :one
UPDATE incomes
SET sub_category_id = $2,
	income_date = $3,
	amount = $4,
	description = $5,
	updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteIncome :one
UPDATE incomes
SET deleted_at = NOW()
WHERE id = $1
RETURNING *;