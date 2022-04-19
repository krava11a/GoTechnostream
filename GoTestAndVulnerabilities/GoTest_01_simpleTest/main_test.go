package main

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	//for parallel testing
	t.Parallel()
	/////////////

	if Sum(1, 3) != 4 {
		t.Error("Error!!", "expected", 4, "got", Sum(1, 3))
	}
}

func TestDivision(t *testing.T) {
	//for parallel testing
	t.Parallel()
	/////////////

	defer func() {
		recover()
	}()

	if Division(2, 1) != 2 {
		t.Errorf("expected %d, got %d", 2, Division(2, 1))
	}

	Division(2, 0)
	t.Error("Panic expected")

}

func TestAddUser(t *testing.T) {
	//for parallel testing
	t.Parallel()
	/////////////

	users := []User{}
	AddUser(&users, "Vasya", 34)

	if len(users) == 0 {
		t.Fatal("Empty slice")
	}

	expected := []User{
		{
			Name: "Vasya",
			Age:  34,
		},
	}

	//for slice and struct we need use reflect.DeepEqual() for compare
	if !reflect.DeepEqual(users, expected) {
		t.Errorf("Expected %+v, got %+v", expected, users)
	}

}

//Table testing удобен когда много входных и выходных значений
//Чтобы не копипастить код сравнения
func TestTable(t *testing.T) {
	//for parallel testing
	t.Parallel()
	time.Sleep(2*time.Second)
	/////////////

	t.Log("Create table with cases")
	cases := []struct {
		A      int
		B      int
		Result int
		Err    error
	}{
		{
			A:      1,
			B:      1,
			Result: 1,
			Err:    nil,
		},
		{
			A:      0,
			B:      1,
			Result: 0,
			Err:    nil,
		},
		{
			A:      1,
			B:      0,
			Result: 0,
			Err:    errors.New("Division by zero"),
		},
	}

	t.Log("Check all cases in cycle")
	for _, testCase := range cases {
		res, err := TrueDivision(testCase.A, testCase.B)
		if res != testCase.Result {
			t.Errorf("Expected %d, got %d", testCase.Result, res)
		}
		if err != nil && err.Error() != testCase.Err.Error() {
			t.Errorf("Expected %+v, got %+v", testCase.Err, err)
		}
	}

}
