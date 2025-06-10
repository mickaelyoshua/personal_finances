-- name: CreateIncome :one
INSERT INTO incomes (user_id, category_id, income_date, amount, description)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetIncome :one
SELECT * FROM incomes WHERE id = $1 LIMIT 1;

-- name: GetAllIncomes :many
SELECT * FROM incomes WHERE user_id = $1 ORDER BY income_date DESC;

-- name: UpdateIncome :one
UPDATE incomes
SET category_id = $2,
	income_date = $3,
	amount = $4,
	description = $5,
	updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteIncome :exec
DELETE FROM incomes
WHERE id = $1;