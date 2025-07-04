// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Expense struct {
	ID            int32
	UserID        int32
	SubCategoryID pgtype.Int4
	ExpenseDate   pgtype.Date
	Amount        pgtype.Numeric
	Description   pgtype.Text
	CreatedAt     pgtype.Timestamptz
	UpdatedAt     pgtype.Timestamptz
}

type ExpenseCategory struct {
	ID   int32
	Name string
}

type ExpenseSubCategory struct {
	ID         int32
	CategoryID int32
	Name       string
}

type Income struct {
	ID          int32
	UserID      int32
	CategoryID  pgtype.Int4
	IncomeDate  pgtype.Date
	Amount      pgtype.Numeric
	Description pgtype.Text
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type IncomeCategory struct {
	ID   int32
	Name string
}

type User struct {
	ID           int32
	Email        string
	Name         string
	PasswordHash string
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
	DeletedAt    pgtype.Timestamptz
}
