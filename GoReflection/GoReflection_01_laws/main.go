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

	type MyInt int
	var c MyInt = 7
	v = reflect.ValueOf(c)
	fmt.Println("kind is int:",v.Kind()==reflect.Int)

	//1 закон рефлексии - перейти к пустому интерфейсу
	y := v.Interface().(MyInt) // y will have type float64
	fmt.Println("Значение обертки",v,"Само значение",y)

	access()
}

func access() {

	//2 закон
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	//изменение v запрещено, тк отсутствует связь с подлежащим значением
	//v.SetFloat(7.1) // will panic

	fmt.Println("settability of v:", v.CanSet())

	//чтобы иметь возможность изменить значение, нам потребуется ссылка
	p:=reflect.ValueOf(&x)
	fmt.Println("type of p:",p.Type())
	fmt.Println("settability of p:",p.CanSet())

	//use elev to change value, that got for pointer
	v = p.Elem()
	fmt.Println("settability of v:",v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)

}