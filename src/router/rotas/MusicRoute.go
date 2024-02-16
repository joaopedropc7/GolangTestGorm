package rotas

import (
	"Routes/src/controllers"
	"net/http"
)

var MusicRoutes = []Rota{
	{
		URI:                "/music",
		Metodo:             http.MethodPost,
		Funcao:             controllers.UploadMusic,
		RequerAutenticacao: true,
	},
	{
		URI:                "/music/{musicid}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetMusicById,
		RequerAutenticacao: true,
	},
}
