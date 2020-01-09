package hpool

type User struct {
	accessKey string
	secretKey string
}

type SubAccount struct {
	user    *User
	SubName string
	SubCode string
}

func NewUser(accessKey, secretKey string) *User {
	return &User{
		accessKey: accessKey,
		secretKey: secretKey,
	}
}

func (u *User) Sub(name, code string) *SubAccount {
	return &SubAccount{
		user:    u,
		SubName: name,
		SubCode: code,
	}
}
