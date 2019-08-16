package hpool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	accessKey := "AAA"
	secretKey := "AAABBBCCC"
	query := "access_key=" + accessKey +
		"&account_name=111c&currency=btc&date=2018-09-11" +
		"&sub_code_list=mnhmn2if&timestamp=1537501071060&uid=1551251"
	sign := Sign(secretKey, query)
	assert.Equal(t, "AAEBA0C4D9936CE7C13D85592D7DB0D77D9F335BDDD186F0888C902E70DD08BE", sign)
}
