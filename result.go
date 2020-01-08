package hpool

import (
	"encoding/json"
	"fmt"
)

type Result interface {
	Error() error
}

type result struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (r *result) Error() error {
	if r.Success {
		return nil
	}
	return fmt.Errorf("HuobiPool %d: %s", r.Code, r.Message)
}

func unmarshal(data []byte, value Result) error {
	err := json.Unmarshal(data, value)
	if err != nil {
		return err
	}
	return value.Error()
}

type TodayProfit struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type WorkerStats struct {
	Currency string `json:"currency"`
	Active   int64  `json:"workers_active,string"`
	Dead     int64  `json:"workers_dead,string"`
	Inactive int64  `json:"workers_inactive,string"`
	Total    int64  `json:"workers_total,string"`
}

type HashRates struct {
	Currency  string      `json:"currency"`
	Speed15m  json.Number `json:"speed_f"`
	Speed1h   json.Number `json:"speed_s"`
	Speed1d   json.Number `json:"speed_t"`
	Unit15m   string      `json:"unit_f"`
	Unit1h    string      `json:"unit_s"`
	Unit1d    string      `json:"unit_t"`
	Reject15m json.Number `json:"reject_f"`
	Reject1h  json.Number `json:"reject_s"`
	Reject1d  json.Number `json:"reject_t"`
}

type TodayProfitResult struct {
	result
	Data TodayProfit `json:"data"`
}

type TodayProfitResultV2 struct {
	result
	Data []TodayProfit `json:"data"`
}

type WorkerStatsResult struct {
	result
	Data WorkerStats `json:"data"`
}

type HashRatesResult struct {
	result
	Data HashRates `json:"data"`
}

type ChangeCoinResult struct {
	result
}
