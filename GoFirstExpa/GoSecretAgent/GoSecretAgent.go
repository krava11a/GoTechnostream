package main

import "fmt"

type Person struct {
	Name string
	inn string
}

type Stuff struct {
	inn int
}

type SecretAgent struct {
	Person
	Stuff
	LicenseToKill bool
}

func (p Person) getName() string  {
	return p.Name
}

func (s SecretAgent) getName() string  {
	return "CLASSIFIELD"
}

func main() {
	sa := SecretAgent{
		Person{"Smith","Zero"},
		Stuff{47},
		true,
	}
	fmt.Println(sa)
	fmt.Println(sa.Stuff.inn)
	fmt.Println(sa.Person.inn," ",sa.Person.Name)
	fmt.Println(sa.getName())
}
