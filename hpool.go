package hpool

import (
	"strconv"
)

// 分页查询子账号的收益明细
func (p *SubAccount) ListRecord(page, size int) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.subCode,
		"page":       strconv.Itoa(page),
		"size":       strconv.Itoa(size),
	}
	request("GET", p.user.secretKey, "/open/api/account/v1/list-record", params)
}

// 查询子账号的实时算力
func (p *SubAccount) GetHashRate() (*HashRates, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.subCode,
	}
	res, err := request("GET", p.user.secretKey,
		"/open/api/user/v1/get-hash-rate", params)
	if err != nil {
		return nil, err
	}
	r := HashRatesResult{}
	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return &r.Data, nil
}

// 查询子账号的矿工统计
func (p *SubAccount) GetWorkerStats() (*WorkerStats, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.subCode,
	}
	res, err := request("GET", p.user.secretKey, "/open/api/user/v1/get-worker-stats", params)
	if err != nil {
		return nil, err
	}
	r := WorkerStatsResult{}
	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return &r.Data, nil
}

// 今日预估收益
func (p *SubAccount) GetTodayProfit() (*TodayProfit, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.subCode,
	}
	res, err := request("GET", p.user.secretKey, "/open/api/user/v1/get-today-profit", params)
	if err != nil {
		return nil, err
	}
	r := TodayProfitResult{}
	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return &r.Data, nil
}

// 今日预估收益V2
func (p *SubAccount) GetTodayProfitV2() ([]TodayProfit, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.subCode,
	}
	res, err := request("GET", p.user.secretKey, "/open/api/user/v2/get-today-profit", params)
	if err != nil {
		return nil, err
	}
	r := TodayProfitResultV2{}
	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

// 切换挖矿币种
func (p *SubAccount) ChangeCoin(coin string) (bool, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.subCode,
		"currency":   coin,
	}
	res, err := request("POST", p.user.secretKey, "/open/api/user/v1/change-sub-user-currency", params)
	if err != nil {
		return false, err
	}
	r := ChangeCoinResult{}
	err = unmarshal(res, &r)
	if err != nil {
		return false, err
	}
	return true, nil
}
