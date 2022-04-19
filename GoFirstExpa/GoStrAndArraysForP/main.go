package main

import "fmt"

func main() {
	s1 := string("Hello")
	s2 := "world"
	s3 := s1+" "+s2
	s4 := s3[2:7]
	println(s3,"   ",s4)



	sa1 := make([]string, 10)
	fmt.Println("sa1 - ",sa1)
	println("sa1 - ",sa1)
	for i:=0;i<len(sa1);i++ {
		sa1[i] = "g"
	}
	//sa2 :=append(sa1[0:])
	sa2 :=append(sa1[0:],[]string{"asd","www"}...)
	sa2[2] = "interconnection"
	//sa1[3] = "intervention"
	fmt.Println("sa2 - ",sa2)
	println("sa2 - ",sa2)
	fmt.Println("sa1 - ",sa1)
	println("sa1 - ",sa1)

}
