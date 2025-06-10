package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"github.com/mickaelyoshua/personal-finances/util"
)

func createRandomIncome() (Income, error) {
	args := CreateIncomeParams{
		UserID:      testUser.ID,
		CategoryID:  util.RandomUUID(),
		IncomeDate:  util.RandomDate(),
		Amount:      util.RandomAmount(),
		Description: util.RandomDescription(),
	}
	return testQueries.CreateIncome(context.Background(), args)
}

func TestCreateIncome(t *testing.T) {
	income, err := createRandomIncome()

	require.NoError(t, err)
	require.NotEmpty(t, income)
}

func TestGetIncome(t *testing.T) {
	income1, err := createRandomIncome()
	require.NoError(t, err)
	require.NotEmpty(t, income1)

	income2, err := testQueries.GetIncome(context.Background(), income1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, income2)

	require.Equal(t, income1.ID, income2.ID)
	require.Equal(t, income1.UserID, income2.UserID)
	require.Equal(t, income1.CategoryID, income2.CategoryID)
	require.Equal(t, income1.Amount, income2.Amount)
	require.Equal(t, income1.Description, income2.Description)

	require.WithinDuration(t, income1.IncomeDate.Time, income2.IncomeDate.Time, time.Second)
	require.WithinDuration(t, income1.CreatedAt.Time, income2.CreatedAt.Time, time.Second)
	require.WithinDuration(t, income1.UpdatedAt.Time, income2.UpdatedAt.Time, time.Second)
}

func TestGetAllIncomes(t *testing.T) {
	numberOfIncomes := 5
	for range numberOfIncomes {
		_, err := createRandomIncome()
		require.NoError(t, err)
	}

	incomes, err := testQueries.GetAllIncomes(context.Background(), testUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, incomes)
	require.GreaterOrEqual(t, len(incomes), numberOfIncomes)

	for _, income := range incomes {
		require.NotEmpty(t, income)
	}
}

func TestUpdateIncome(t *testing.T) {
	income1, err := createRandomIncome()
	require.NoError(t, err)
	require.NotEmpty(t, income1)

	args := UpdateIncomeParams{
		ID:          income1.ID,
		CategoryID: income1.CategoryID, // Keeping the same subcategory for simplicity
		IncomeDate:  income1.IncomeDate,      // Keeping the same date for simplicity
		Amount:     util.RandomAmount(),
		Description: util.RandomDescription(),
	}

	income2, err := testQueries.UpdateIncome(context.Background(), args)
	updateTime := time.Now()

	require.NoError(t, err)
	require.NotEmpty(t, income2)

	require.NotEqual(t, income1.Amount, income2.Amount)	 // Amount should change
	require.NotEqual(t, income1.Description, income2.Description)   // Description should change
	require.NotEqual(t, income1.UpdatedAt.Time, income2.UpdatedAt.Time) // UpdatedAt should change
	require.WithinDuration(t, updateTime, income2.UpdatedAt.Time, 2*time.Second) // UpdatedAt should be close to now

	require.Equal(t, income1.ID, income2.ID)
	require.Equal(t, income1.UserID, income2.UserID)
	require.Equal(t, income1.CategoryID, income2.CategoryID)
	require.Equal(t, args.Amount, income2.Amount)
	require.Equal(t, args.Description, income2.Description)
	require.WithinDuration(t, income1.IncomeDate.Time, income2.IncomeDate.Time, 0)
}

func TestDeleteIncome(t *testing.T) {
	income, err := createRandomIncome()
	require.NoError(t, err)
	require.NotEmpty(t, income)

	err = testQueries.DeleteIncome(context.Background(), income.ID)
	require.NoError(t, err)

	deletedIncome, err := testQueries.GetIncome(context.Background(), income.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, deletedIncome)
}