package hpool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	AccessKey = "your_access_key"
	SecretKey = "your_secret_key"
	SubCode   = "your_sub_account_code"
)

func TestGetTodayProfit(t *testing.T) {
	account := NewUser(AccessKey, SecretKey).Sub(SubCode)
	data, err := account.GetTodayProfit()
	assert.NoError(t, err)
	assert.True(t, data.Amount >= 0)
}

func TestGetTodayProfitV2(t *testing.T) {
	account := NewUser(AccessKey, SecretKey).Sub(SubCode)
	data, err := account.GetTodayProfitV2()
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}

func TestChangeCoin(t *testing.T) {
	account := NewUser(AccessKey, SecretKey).Sub(SubCode)
	success, err := account.ChangeCoin("bch")
	assert.Nil(t, err)
	assert.True(t, success)
}
