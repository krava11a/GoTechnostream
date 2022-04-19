package main

import (
	"fmt"
)

type Flyer interface {
	Fly()
}

type Bird struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Println(b.Name + " is flying!!!")
}

func DoFly(f Flyer) {
	f.Fly()
}

type Mig45 struct {
	Name string
}

func (m Mig45) Fly() {
	fmt.Println(m.Name + " is flying!!!! WZHUHHH")
}

func GoFly(f Flyer) {
	f.Fly()
	//b := f.(Bird)
	if b, ok := f.(Bird); ok {
		fmt.Println(b.Name)
	}else if m,ok := f.(Mig45); ok{
		fmt.Println(m.Name)
	}
}




func main() {
	b := Bird{Name: "Cryacva"}
	m := Mig45{Name: "MIG 45 "}
	DoFly(b)
	DoFly(m)
	f := b
	f.Fly()
	GoFly(b)
	GoFly(m)





}
