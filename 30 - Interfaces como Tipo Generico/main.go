package main

import "fmt"

func gererica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	gererica("String")
	gererica(1)
	gererica(true)

	fmt.Println(1, 2, "String", false, true, float64(1.23456789101))
}
