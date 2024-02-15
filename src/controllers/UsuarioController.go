package controllers

import (
	"Routes/src/models"
	"Routes/src/response"
	"Routes/src/service"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var usuarioInput models.UsuarioRequestVO
	if err := json.NewDecoder(r.Body).Decode(&usuarioInput); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	usuario, err := service.CreateUsuarioService(usuarioInput)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusOK, usuario)
}
