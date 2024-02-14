package banco

import (
	"Routes/src/config"
	"Routes/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Conectar() (*gorm.DB, error) {
	dsn := config.StringConexaoBanco

	novoDB, erro := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if erro != nil {
		return nil, erro
	}

	db = novoDB

	if err := db.AutoMigrate(&models.Product{}); err != nil {
		return nil, err
	}

	return db, nil
}
