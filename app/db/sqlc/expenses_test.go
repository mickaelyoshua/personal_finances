package sqlc

import (
	"time"
	"testing"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"github.com/mickaelyoshua/personal-finances/util"
)

func deleteRandomExpense(expenseID int32) error {
	return testQueries.DeleteExpense(context.Background(), expenseID)
}

func createRandomExpense() (Expense, error){
	args := CreateExpenseParams{
		UserID:    testUser.ID,
		SubCategoryID: util.RandomUUID(),
		ExpenseDate: util.RandomDate(),
		Amount:    util.RandomAmount(),
		Description: util.RandomDescription(),
	}
	return testQueries.CreateExpense(context.Background(), args)
}

func TestCreateExpense(t *testing.T) {
	expense, err := createRandomExpense()
	// delete the expense after test
	defer func() {
		err = deleteRandomExpense(expense.ID)
		require.NoError(t, err)
	}()

	require.NoError(t, err)
	require.NotEmpty(t, expense)
}

func TestGetExpense(t *testing.T) {
	expense1, err := createRandomExpense()
	require.NoError(t, err)
	require.NotEmpty(t, expense1)
	// delete the expense after test
	defer func() {
		err = deleteRandomExpense(expense1.ID)
		require.NoError(t, err)
	}()

	expense2, err := testQueries.GetExpense(context.Background(), expense1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, expense2)

	require.Equal(t, expense1.ID, expense2.ID)
	require.Equal(t, expense1.UserID, expense2.UserID)
	require.Equal(t, expense1.SubCategoryID, expense2.SubCategoryID)
	require.Equal(t, expense1.Amount, expense2.Amount)
	require.Equal(t, expense1.Description, expense2.Description)

	require.WithinDuration(t, expense1.ExpenseDate.Time, expense2.ExpenseDate.Time, time.Second)
	require.WithinDuration(t, expense1.CreatedAt.Time, expense2.CreatedAt.Time, time.Second)
	require.WithinDuration(t, expense1.UpdatedAt.Time, expense2.UpdatedAt.Time, time.Second)
}

func TestGetAllExpenses(t *testing.T) {
	var expensesID []int32
	numberOfExpenses := 5
	for range numberOfExpenses {
		expense, err := createRandomExpense()
		require.NoError(t, err)
		expensesID = append(expensesID, expense.ID)
	}
	// delete expenses after test
	defer func() {
		for _, expenseID := range expensesID {
			err := deleteRandomExpense(expenseID)
			require.NoError(t, err)
		}
	}()

	expenses, err := testQueries.GetAllExpenses(context.Background(), testUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, expenses)
	require.GreaterOrEqual(t, len(expenses), numberOfExpenses)

	for _, expense := range expenses {
		require.NotEmpty(t, expense)
	}
}

func TestUpdateExpense(t *testing.T) {
	expense1, err := createRandomExpense()
	require.NoError(t, err)
	require.NotEmpty(t, expense1)
	// delete the expense after test
	defer func() {
		err = deleteRandomExpense(expense1.ID)
		require.NoError(t, err)
	}()

	args := UpdateExpenseParams{
		ID:          expense1.ID,
		SubCategoryID: expense1.SubCategoryID, // Keeping the same subcategory for simplicity
		ExpenseDate: expense1.ExpenseDate, // Keeping the same date for simplicity
		Amount:     util.RandomAmount(),
		Description: util.RandomDescription(),
	}

	expense2, err := testQueries.UpdateExpense(context.Background(), args)
	updateTime := time.Now()

	require.NoError(t, err)
	require.NotEmpty(t, expense2)

	require.NotEqual(t, expense1.Amount, expense2.Amount) // Amount should change
	require.NotEqual(t, expense1.Description, expense2.Description) // Description should change
	require.NotEqual(t, expense1.UpdatedAt.Time, expense2.UpdatedAt.Time) // UpdatedAt should change
	require.WithinDuration(t, updateTime, expense2.UpdatedAt.Time, 2*time.Second) // UpdatedAt should be close to now

	require.Equal(t, expense1.ID, expense2.ID)
	require.Equal(t, expense1.UserID, expense2.UserID)
	require.Equal(t, expense1.SubCategoryID, expense2.SubCategoryID)
	require.Equal(t, args.Amount, expense2.Amount)
	require.Equal(t, args.Description, expense2.Description)
	require.WithinDuration(t, expense1.ExpenseDate.Time, expense2.ExpenseDate.Time, 0)
}

func TestDeleteExpense(t *testing.T) {
	expense, err := createRandomExpense()
	require.NoError(t, err)
	require.NotEmpty(t, expense)

	err = testQueries.DeleteExpense(context.Background(), expense.ID)
	require.NoError(t, err)

	deletedExpense, err := testQueries.GetExpense(context.Background(), expense.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, deletedExpense)
}