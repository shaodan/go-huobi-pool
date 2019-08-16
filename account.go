package hpool

type HPUser struct {
	accessKey string
	secretKey string
}

type HPSubAccount struct {
	user    *HPUser
	subCode string
}
