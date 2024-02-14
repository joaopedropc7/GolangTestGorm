package main

import (
	"Routes/src/banco"
	"Routes/src/config"
	"Routes/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	fmt.Println(config.Porta)

	fmt.Println("Rodando API!")
	fmt.Println(config.SecretKey)

	db, err := banco.Conectar()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	var result int
	if err := db.Raw("SELECT 1").Scan(&result).Error; err != nil {
		log.Fatal("Erro ao testar a conexão com o banco de dados:", err)
	}

	if err != nil {
		log.Fatal("Erro ao testar a conexão com o banco de dados:", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
