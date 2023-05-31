package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var arg CreateEntryParams

func createTestEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	arg = CreateEntryParams{
		AccountID: account.ID,
		Amount: 10,
	}
	newEntry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)

	return newEntry
}

func TestCreateEntry(t *testing.T)  {
	entry := createTestEntry(t)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
}

func TestGetEntry(t *testing.T)  {
	entry := createTestEntry(t)
	queryEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, queryEntry)

	require.Equal(t, entry.ID, queryEntry.ID)
	require.Equal(t, entry.AccountID, queryEntry.AccountID)
	require.Equal(t, entry.Amount, queryEntry.Amount)
	require.WithinDuration(t, entry.CreatedAt, queryEntry.CreatedAt, time.Second)
}