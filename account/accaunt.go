package account

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updateAt"`
}

func (acc *Account) OutputData() {
	color.Red("%v, %v, %v", acc.Login, acc.Password, acc.Url)
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
	acc.Password = string(res)
}

func (acc *Account) ToBytes() ([]byte, error) {
	return json.Marshal(acc)
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
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
		Login:     login,
		Password:  password,
		Url:       urlData,
	}
	// field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	// fmt.Println(field.Tag)
	// fmt.Println(string(field.Tag))
	if password == "" {
		newAcc.generatePassword(8)
	}
	return newAcc, nil
}
