package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
func main() {
	//Concorrência != Paralelismo
	go escrever("Olá Mundo!")
	escrever("Programando em Go!")
}
