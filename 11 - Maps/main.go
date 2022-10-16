package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"Curso": {
			"Nome":   "Engenharia",
			"Campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	usuario2["Signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)
}
