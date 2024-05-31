package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticação: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticação: false,
	},
	{
		URI:                "/usuarios/{usuariosId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticação: false,
	},
	{
		URI:                "/usuarios/{usuariosId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticação: false,
	},
	{
		URI:                "/usuarios/{usuariosId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticação: false,
	},
}
