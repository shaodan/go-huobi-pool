package hpool

import (
	"encoding/json"
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	client = resty.New().
		SetHostURL("https://openapi.hpt.com").
		SetTimeout(5 * time.Second)
)

func encodeQuery(params map[string]string) string {
	var values []string
	for k := range params {
		values = append(values, k)
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	var sb strings.Builder
	first := true
	for _, k := range values {
		if first {
			first = false
		} else {
			sb.WriteString("&")
		}
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(params[k])
	}
	return sb.String()
}

func request(method, secretKey, path string, params map[string]string) (res []byte, err error) {
	ts := time.Now().UnixNano() / 1e6
	params["timestamp"] = strconv.FormatInt(ts, 10)
	query := encodeQuery(params)
	sign := Sign(secretKey, query)
	req := client.R()
	var resp *resty.Response
	switch method {
	case "GET":
		fullQuery := query + "&sign=" + sign
		resp, err = req.SetQueryString(fullQuery).Get(path)
	case "POST":
		params["sign"] = sign
		var body []byte
		body, err = json.Marshal(params)
		if err != nil {
			err = errors.New("marshal params error" + err.Error())
		} else {
			resp, err = req.SetHeader("Content-Type", "application/json").
				SetBody(body).
				Post(path)
		}
	default:
		err = errors.New("unmatched http method: " + method)
	}
	if err != nil {
		return
	}
	res = resp.Body()
	return
}
