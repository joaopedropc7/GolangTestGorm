package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	StringConexaoBanco = ""

	//Porta onde a API vai estar rodando
	Porta = 0

	// SecretKey é a chave que vai ser usada para ssinar o token
	SecretKey []byte
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s sslmode=disable",
		os.Getenv("DB_Usuario"),
		os.Getenv("DB_NOME"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_HOST"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
