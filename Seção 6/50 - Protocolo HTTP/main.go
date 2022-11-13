package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func Usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}

func main() {
	//HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DE COMUNICAÇÃO WEB

	// CLIENTE(FAZ A REQUISITCAO) - SERVIDOR(PROCESSA A REQUISIÇÃO E ENVIA A RESPOSTA)

	//Request - Response

	//Rotas
	//URI - Identificador do Recurso
	//Método - GET,POST,PUT,DELETE

	http.HandleFunc("/home", home)

	http.HandleFunc("/Usuarios", Usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil)) //Servidor já está rodando
}
