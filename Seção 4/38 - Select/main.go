package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string, 10), make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 500)
		canal1 <- "Canal 1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		canal2 <- "Canal 2"
	}()

	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)

		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}

	}

}
