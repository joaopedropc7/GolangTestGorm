package rotas

import (
	"Routes/src/controllers"
	"net/http"
)

var UsuarioRoutes = []Rota{
	{
		URI:                "/user",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/login",
		Metodo:             http.MethodPost,
		Funcao:             controllers.Login,
		RequerAutenticacao: false,
	},
}
