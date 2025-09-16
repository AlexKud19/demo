package main

import (
	"demo/password/account"
	"fmt"
)

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(myAccount)
	myAccount.OutputData()
}

func promptData(text string) string {
	fmt.Print(text, ": ")
	var value string
	fmt.Scanln(&value)
	return value
}

// func reverce(array *[4]int) {
// 	length := len(*array)
// 	for index, value := range *array {
// 		(*array)[length-index-1] = value
// 	}
// }

// func reverce2(array *[4]int) {
// 	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
// 		array[i], array[j] = array[j], array[i]
// 	}
// }
