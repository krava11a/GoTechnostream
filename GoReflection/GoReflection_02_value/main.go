package main

import (
	"fmt"
	"reflect"
)

func printAny(arg interface{}) {
	printValue(reflect.ValueOf(arg))
	fmt.Print("\n")
}

func printValue(v reflect.Value) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("int %v", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Printf("Uint %v", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Printf("float %v", v.Float())

	case reflect.String:
		fmt.Printf("string %v",v.String())

	case reflect.Bool:
		fmt.Printf("bool %v",v.Bool())
	case reflect.Map:
		fmt.Print("map{")
		for i, key := range v.MapKeys() {
			if i > 0 {
				fmt.Print(", ")
			}
			printValue(key)
			fmt.Print(":")
			printValue(v.MapIndex(key))
		}
		fmt.Print("}")
	case reflect.Interface, reflect.Ptr:
		fmt.Print("pointer ")
		printValue(v.Elem())

	}


}

func main() {
	printAny("Arte")
	printAny(12)


	m := make(map[string]int)
	m["first"] = 1
	m["second"] = 3
	m["third"] = 2
	printAny(m)


	p := &m
	printAny(p)


}
