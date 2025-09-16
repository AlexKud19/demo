package account

import (
	"fmt"
	"math/rand/v2"
	"net/url"

	"github.com/fatih/color"
)

type Account struct {
	login    string
	password string
	url      string
}

func (acc *Account) OutputData() {
	color.Red("%v, %v, %v", acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	chars := make([]rune, 0, 94)
	for i := rune(33); i < 126; i++ {
		chars = append(chars, i)
	}
	for i := range n {
		res[i] = chars[rand.IntN(93)]
	}
	acc.password = string(res)
}

func NewAccount(login, password, urlData string) (*Account, error) {
	if login == "" {
		return nil, fmt.Errorf("empty login")
	}
	_, err := url.ParseRequestURI(urlData)
	if err != nil {
		return nil, fmt.Errorf("invalid url")
	}
	newAcc := &Account{
		login:    login,
		password: password,
		url:      urlData,
	}
	if password == "" {
		newAcc.generatePassword(8)
	}
	return newAcc, nil
}
