package repository

import (
	"Routes/src/models"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	DB *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{db}
}

func (r *UsuarioRepository) CreateUsuario(Usuario *models.Usuario) (*models.Usuario, error) {
	if err := r.DB.Create(Usuario).Error; err != nil {
		return nil, err
	}
	return Usuario, nil
}
