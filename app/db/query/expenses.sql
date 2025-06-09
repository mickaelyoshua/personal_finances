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

-- name: DeleteExpense :exec
UPDATE expenses
SET deleted_at = NOW()
WHERE id = $1;