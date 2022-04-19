package main

import (
	"fmt"
	"reflect"
)

func main() {
	x:= 3.4
	//TypeOf() принимает на вход interface{}, в этом месте будет аллокация
	fmt.Println("reflect.Type:", reflect.TypeOf(x))
	fmt.Printf("reflect.Type: %+v\n", reflect.TypeOf(x))

	//relect.Value != значению переданному на вход
	fmt.Println("reflect.Value:", reflect.ValueOf(x).String())

	v:= reflect.ValueOf(x)
	fmt.Println("Тип value: ",v.Type())
	//kind получение базового типа
	fmt.Println("Тип float64: ",v.Kind() == reflect.Float64)
	fmt.Println("Тип value: ",v.Float())
}
