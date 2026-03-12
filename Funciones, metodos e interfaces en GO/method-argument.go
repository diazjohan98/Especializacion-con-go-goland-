package main

import "fmt"

type Myint int

func (mi Myint) Double() int {
	return int(mi * 2)
}

func (mi Myint) Triple() int {
	return int(mi * 3)
}

func main() {
	v := Myint(3)
	fmt.Println(v.Double())
	fmt.Println(v.Triple())
}
