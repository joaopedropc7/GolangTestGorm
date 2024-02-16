package repository

import (
	"Routes/src/models"
	"errors"
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

func (r *UsuarioRepository) FindUserByEmail(email string) (models.Usuario, error) {
	var usuario models.Usuario
	err := r.DB.Where("email = ?", email).First(&usuario).Error
	if err != nil {
		return models.Usuario{}, errors.New("Não foi encontrado nenhum registro!")
	}
	return usuario, err
}

func (r *UsuarioRepository) GetUserById(usuarioId uint64) (models.Usuario, error) {
	var usuario models.Usuario
	if err := r.DB.First(&usuario, usuarioId).Error; err != nil {
		return models.Usuario{}, errors.New("não foi encontrado nenhum registro com este ID")
	}
	return usuario, nil
}
