package middlewares

import (
	"Routes/src/authentication"
	"Routes/src/response"
	"log"
	"net/http"
)

// Escreve informacoes da requisicao no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// AUTENTICAR VERIFICA SE O Usuario FAZENDO A REQUISICAO ESTA AUTENTICADO
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := authentication.ValidateToken(r); erro != nil {
			response.Error(w, http.StatusUnauthorized, erro)
			return
		}
		proximaFuncao(w, r)
	}
}
