package sqlc

import (
	"time"
	"testing"
	"context"


	"github.com/stretchr/testify/require"
	"github.com/mickaelyoshua/personal-finances/util"
)

func deleteRandomExpense(expenseID int32) error {
	return testQueries.DeleteExpense(context.Background(), expenseID)
}

func createRandomExpense() (Expense, error){
	user, err := createRandomUser()
	if err != nil {
		return Expense{}, err
	}

	args := CreateExpenseParams{
		UserID:    user.ID,
		SubCategoryID: util.RandomUUID(),
		ExpenseDate: util.RandomDate(),
		Amount:    util.RandomAmount(),
		Description: util.RandomDescription(),
	}
	return testQueries.CreateExpense(context.Background(), args)
}

func TestCreateExpense(t *testing.T) {
	expense, err := createRandomExpense()

	require.NoError(t, err)
	require.NotEmpty(t, expense)

	// delete the expense after test
	err = deleteRandomExpense(expense.ID)
	require.NoError(t, err)
}

func TestGetExpense(t *testing.T) {
	expense1, err := createRandomExpense()
	require.NoError(t, err)
	require.NotEmpty(t, expense1)

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

	// delete the expense after test
	err = deleteRandomExpense(expense1.ID)
	require.NoError(t, err)
}