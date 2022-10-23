package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) Salvar() { //Método de usuario
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados \n", u.nome)
}

func (u usuario) MaiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) FazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 17}
	fmt.Println(usuario1)
	usuario1.Salvar()

	MaiorIdade := usuario1.MaiorDeIdade()
	fmt.Println(MaiorIdade)

	usuario1.FazerAniversario()

	MaiorIdade2 := usuario1.MaiorDeIdade()
	fmt.Println(MaiorIdade2)

}
