package rotas

import (
	"Routes/src/controllers"
	"net/http"
)

var productRoutes = []Rota{
	{
		URI:                "/product",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateProduct,
		RequerAutenticacao: false,
	},
	{
		URI:                "/product/{productid}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindProductById,
		RequerAutenticacao: false,
	},
	{
		URI:                "/product",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindAllProducts,
		RequerAutenticacao: false,
	},
	{
		URI:                "/product/{productid}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateProductParsingId,
		RequerAutenticacao: false,
	},
	{
		URI:                "/product/{productid}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteProductParsingId,
		RequerAutenticacao: false,
	},
}
