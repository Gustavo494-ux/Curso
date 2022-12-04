package rotas

import (
	"api/src/router/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/Usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/Usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:                "/Usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/Usuarios",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/Usuarios",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
