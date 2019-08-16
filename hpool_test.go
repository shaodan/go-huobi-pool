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
	account := New(AccessKey, SecretKey).Sub(SubCode)
	data, err := account.GetTodayProfit()
	assert.Nil(t, err)
	assert.True(t, data.Amount >= 0)
}

func TestChangeCoin(t *testing.T) {
	account := New(AccessKey, SecretKey).Sub(SubCode)
	success, err := account.ChangeCoin("bch")
	assert.Nil(t, err)
	assert.True(t, success)
}
