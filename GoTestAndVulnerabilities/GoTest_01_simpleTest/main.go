package main

import "errors"

func main() {

}

func Sum(a, b int) int {
	return a + b
}

func Division(a, b int) int {
	return a / b
}

func TrueDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}

type User struct {
	Name string
	Age  int
}

func AddUser(users *[]User, name string, age int) {
	u := User{
		Name: name,
		Age:  age,
	}

	*users = append(*users, u)
	return
}