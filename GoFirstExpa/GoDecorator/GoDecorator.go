package main

import "fmt"

type isStuff interface {
	DoStuff()
}

type realStuff string

func (r realStuff) DoStuff()  {
	fmt.Println(r)
}

type fakeStuff int

func (f fakeStuff) DoStuff()  {
	fmt.Println("It is fakeStuff")
}

type stuff struct {
	isStuff
	Name string
}

func (s stuff) SomeComplex()  {
	s.DoStuff()
}

func main() {
	r := realStuff("Hi")
	f := fakeStuff(0)

	rs := stuff{r,"real stuff"}
	rs.DoStuff()

	fs := stuff{f,"fake stuff"}
	fs.DoStuff()

}
