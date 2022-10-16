package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Domingo"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough //Faz com que o código da proxima condição seja executado.
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana(2))
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana(4))
	fmt.Println(diaDaSemana(5))
	fmt.Println(diaDaSemana(6))
	fmt.Println(diaDaSemana(7))
	fmt.Println(diaDaSemana(8))

	fmt.Println("-------------------------------------")

	fmt.Println(diaDaSemana2(1))
	fmt.Println(diaDaSemana2(2))
	fmt.Println(diaDaSemana2(3))
	fmt.Println(diaDaSemana2(4))
	fmt.Println(diaDaSemana2(5))
	fmt.Println(diaDaSemana2(6))
	fmt.Println(diaDaSemana2(7))
	fmt.Println(diaDaSemana2(8))

	fmt.Println("-------------------------------------")
	fmt.Println(diaDaSemana2(1))

}
