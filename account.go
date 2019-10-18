package hpool

type HPUser struct {
	accessKey string
	secretKey string
}

type HPSubAccount struct {
	user    *HPUser
	subCode string
}

func NewUser(accessKey, secretKey string) *HPUser {
	return &HPUser{
		accessKey: accessKey,
		secretKey: secretKey,
	}
}

func (u *HPUser) Sub(subCode string) *HPSubAccount {
	return &HPSubAccount{
		user:    u,
		subCode: subCode,
	}
}
