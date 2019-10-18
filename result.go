package hpool

import (
	"encoding/json"
	"fmt"
)

type result struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (r *result) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, r)
	if err != nil {
		return err
	}
	if r.Success {
		return nil
	}
	return fmt.Errorf("HuobiPool error %d: %s", r.Code, r.Message)
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

type ChangeCoinResult struct {
	result
}
