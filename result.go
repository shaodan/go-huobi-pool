package hpool

type result struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TodayProfit struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type TodayProfitResult struct {
	result
	Data TodayProfit `json:"data"`
}

type TodayProfitResultV2 struct {
	result
	Data []TodayProfit `json:"data"`
}

type ChangeCoinResult struct {
	result
}
