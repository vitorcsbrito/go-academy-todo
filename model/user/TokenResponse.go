package user

import "time"

type Token struct {
	Token   string        `json:"token"`
	Age     time.Duration `json:"age"`
	Expires time.Time     `json:"expires"`
}

func NewToken(token string, age time.Duration, expires time.Time) *Token {
	return &Token{
		Token:   token,
		Age:     age,
		Expires: expires,
	}
}
