package db

import (
	"context"
	"database/sql"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner: user.Username,
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateAccount(t *testing.T)  {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T)  {
	createdAccount := createRandomAccount(t)
	queryAccount,err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, queryAccount)

	require.Equal(t, createdAccount.ID, queryAccount.ID)
	require.Equal(t, createdAccount.Owner, queryAccount.Owner)
	require.Equal(t, createdAccount.Balance, queryAccount.Balance)
	require.Equal(t, createdAccount.Currency, queryAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, queryAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T)  {
	createdAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID: createdAccount.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, updatedAccount.ID, createdAccount.ID)
	require.Equal(t, updatedAccount.Owner, createdAccount.Owner)
	require.Equal(t, updatedAccount.Balance, arg.Balance)
	require.Equal(t, updatedAccount.Currency, createdAccount.Currency)
}

func TestDeleteAccount(t *testing.T)  {
	createdAccount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	
	searchAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, searchAccount)
}

func TestListAccounts(t *testing.T)  {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner: lastAccount.Owner,
		Limit: 5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}