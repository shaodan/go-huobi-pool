package hpool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	accessKey = "your_access_key"
	secretKey = "your_secret_key"
	subName   = "your_sub_account_name/anything"
	subCode   = "your_sub_account_code"
)

func TestGetTodayProfit(t *testing.T) {
	account := NewUser(accessKey, secretKey).Sub(subName, subCode)
	data, err := account.GetTodayProfit()
	assert.NoError(t, err)
	assert.True(t, data.Amount >= 0)
}

func TestGetTodayProfitV2(t *testing.T) {
	account := NewUser(accessKey, secretKey).Sub(subName, subCode)
	data, err := account.GetTodayProfitV2()
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}

func TestGetWorkerStats(t *testing.T) {
	account := NewUser(accessKey, secretKey).Sub(subName, subCode)
	data, err := account.GetWorkerStats("btc")
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Currency)
}

func TestGetHashRates(t *testing.T) {
	account := NewUser(accessKey, secretKey).Sub(subName, subCode)
	data, err := account.GetHashRate()
	assert.NoError(t, err)
	if assert.NotEmpty(t, data) {
		assert.GreaterOrEqual(t, data.Speed15m, 0.0)
		assert.GreaterOrEqual(t, data.Speed1h, 0.0)
		assert.GreaterOrEqual(t, data.Speed1d, 0.0)
		assert.GreaterOrEqual(t, data.Reject15m, 0.0)
		assert.Less(t, data.Reject15m, 0.1)
	}
}

func TestGetWorkers(t *testing.T) {
	account := NewUser(accessKey, secretKey).Sub(subName, subCode)
	data, err := account.GetWorkers()
	assert.NoError(t, err)
	if assert.NotEmpty(t, data) {
		if assert.NotEmpty(t, data.List) {
			for _, worker := range data.List {
				assert.Equal(t, worker.Status, 1)
				assert.Greater(t, worker.Speed15m, 0.0)
				assert.Less(t, worker.Reject15m, 0.1)
				assert.Greater(t, worker.LastShare, int64(0))
			}
		}
	}
}

func TestChangeCoin(t *testing.T) {
	account := NewUser(accessKey, secretKey).Sub(subName, subCode)
	success, err := account.ChangeCoin("bch")
	assert.Nil(t, err)
	assert.True(t, success)
}

func TestGetTransferProfit(t *testing.T) {
	account := NewUser(accessKey, secretKey).Sub(subName, subCode)
	data, err := account.GetTransferProfit("btc", "2020-02-08")
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}

func TestGetUnitCurrencyProfits(t *testing.T) {
	account := NewUser(accessKey, secretKey).Sub(subName, subCode)
	data, err := account.GetUnitCurrencyProfits("2020-03-10", "btc")
	if assert.NoError(t, err) {
		assert.NotEmpty(t, data)
	}
}
