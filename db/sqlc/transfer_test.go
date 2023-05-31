package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var args CreateTransferParams

func createTestTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	args = CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: 10,
	}
	newTransfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)

	return newTransfer
}

func TestCreateTransfer(t *testing.T)  {
	transfer := createTestTransfer(t)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, args.FromAccountID)
	require.Equal(t, transfer.ToAccountID, args.ToAccountID)
	require.Equal(t, transfer.Amount, args.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}

func TestGetTransfer(t *testing.T)  {
	transfer := createTestTransfer(t)
	queryTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, queryTransfer)

	require.Equal(t, transfer.ID, queryTransfer.ID)
	require.Equal(t, transfer.FromAccountID, queryTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, queryTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, queryTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, queryTransfer.CreatedAt, time.Second)
}