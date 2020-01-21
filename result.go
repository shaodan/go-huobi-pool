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
	Currency  string  `json:"currency"`
	Speed15m  float64 `json:"speed_f,string"`
	Speed1h   float64 `json:"speed_s,string"`
	Speed1d   float64 `json:"speed_t,string"`
	Unit15m   string  `json:"unit_f"`
	Unit1h    string  `json:"unit_s"`
	Unit1d    string  `json:"unit_t"`
	Reject15m float64 `json:"reject_f,string"`
	Reject1h  float64 `json:"reject_s,string"`
	Reject1d  float64 `json:"reject_t,string"`
}

type WorkerList struct {
	List []struct {
		Name      string  `json:"worker_name"`
		Currency  string  `json:"currency"`
		Speed15m  float64 `json:"hash_rate_15m,string"`
		Speed1d   float64 `json:"hash_rate_1d,string"`
		Reject15m float64 `json:"reject15m,string"`
		Reject1d  float64 `json:"reject1d,string"`
		LastShare int64   `json:"last_share_time"`
		Status    int     `json:"status"` // [0:不活跃;1:活跃;2.失效]
	}
	Active     int `json:"workers_active"`
	Inactive   int `json:"workers_inactive"`
	Pagination struct {
		Total int `json:"total_count"`
	} `json:"pagenation"`
}

type TransferProfit struct {
	Speed     float64 `json:"speed,string"`
	Amount    float64 `json:"transfer_amount,string"`
	HPTAmount float64 `json:"hpt_transfer_amount,string"`
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

type WorkersResult struct {
	result
	Data WorkerList `json:"data"`
}

type TransferProfitResult struct {
	result
	Data TransferProfit `json:"data"`
}
