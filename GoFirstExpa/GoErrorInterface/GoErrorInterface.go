package main

import (
	"errors"
	"fmt"
)

// Create a named type for our new error type.
type errorString string

//Implement the error interface
func (e errorString) Error() string {
	return string(e)
}

//New creates interface values of type error.
func New(text string) error {
	return errorString(text)
}

var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

//type error interface {
//	Error() string
//}

type MyError struct {
	Msg  string
	File string
	Line int
}

func (m *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", m.File, m.Line, m.Msg)
}

func BadFunc() error {
	return &MyError{"Something happened", "server.go", 42}
}

func DoIfaces(slice ...interface{}) error {
	return nil
}

func main() {
	if ErrNamedType == New("EOF") {
		fmt.Println("Named Type Error")
	}

	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")
	}

	err := BadFunc()

	switch err := err.(type) {
	case nil:
	//call succeeded, nothing to do
	case *MyError:
		fmt.Println("error occured on line:", err.Line)
	default:
		//unknown error
	}

	st := []int{15,23,43}
	ist := make([]interface{},len(st))
	for i:= range st {
		ist[i] = st[i]
	}
	//copy(ist,st) // not valid
	DoIfaces(ist...)


}
