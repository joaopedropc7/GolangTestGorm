package controllers

import (
	"Routes/src/authentication"
	"Routes/src/models"
	"Routes/src/response"
	"Routes/src/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func UploadMusic(w http.ResponseWriter, r *http.Request) {

	usuarioID, errExtract := authentication.ExtratUsuarioID(r)
	if errExtract != nil {
		response.Error(w, http.StatusUnauthorized, errExtract)
		return
	}

	usuarioInDB, errServiceGetUserInDB := service.FindUserById(usuarioID)
	if errServiceGetUserInDB != nil {
		response.Error(w, http.StatusInternalServerError, errServiceGetUserInDB)
		return
	}

	jsonData := r.FormValue("json_field")

	var musicInput models.MusicRequestVO
	erro := json.Unmarshal([]byte(jsonData), &musicInput)
	if erro != nil {
		http.Error(w, "Erro ao decodificar o JSON", http.StatusBadRequest)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // Limitar o tamanho do arquivo para 10 MB
	if err != nil {
		http.Error(w, "Erro ao analisar o formulário", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erro ao receber o arquivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Ler os dados do arquivo
	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Erro ao ler o conteúdo do arquivo", http.StatusInternalServerError)
		return
	}

	createdMusic, err := service.CreateMusic(&musicInput, fileData, usuarioInDB)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao criar a música: %v", err), http.StatusInternalServerError)
		return
	}

	response.JSON(w, http.StatusOK, createdMusic)
}

func GetMusicById(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	musicId, err := strconv.ParseInt(parametros["musicid"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	music, err := service.FindMusicById(musicId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	file, err := os.Open(music.Path)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "audio/mpeg")

	http.ServeContent(w, r, music.Title, time.Now(), file)
}
