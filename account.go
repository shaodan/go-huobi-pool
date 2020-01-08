package hpool

type User struct {
	accessKey string
	secretKey string
}

type SubAccount struct {
	user    *User
	subCode string
}

func NewUser(accessKey, secretKey string) *User {
	return &User{
		accessKey: accessKey,
		secretKey: secretKey,
	}
}

func (u *User) Sub(subCode string) *SubAccount {
	return &SubAccount{
		user:    u,
		subCode: subCode,
	}
}
