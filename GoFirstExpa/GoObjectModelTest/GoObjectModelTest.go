package main

import "fmt"


type MyInt int

func (m MyInt) showYourSelf() {
	fmt.Printf("%T  %v  \n",m,m)

}

func (m *MyInt) add(i MyInt)  {
	*m = *m + MyInt(i)
	zss := &m
	fmt.Println(zss)
}

func (m MyInt) addWithoutAsterisk(i MyInt)  {
	m = m +i
	fmt.Println("in function : m=",m)
}

func main() {
	var z MyInt = 7
	z.showYourSelf()
	z.add(4)
	z.showYourSelf()
	z.addWithoutAsterisk(455)
	z.showYourSelf()

}

