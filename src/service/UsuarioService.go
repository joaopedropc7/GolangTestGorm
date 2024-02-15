package service

import (
	"Routes/src/banco"
	"Routes/src/models"
	"Routes/src/repository"
)

func CreateUsuarioService(usuarioRequest models.UsuarioRequestVO) (*models.Usuario, error) {
	usuario := &models.Usuario{
		Nome:  usuarioRequest.Nome,
		Nick:  usuarioRequest.Nick,
		Email: usuarioRequest.Email,
		Senha: usuarioRequest.Senha,
	}

	if erro := usuario.Preparar("cadastro"); erro != nil {
		return nil, erro
	}

	db, err := banco.Conectar()
	if err != nil {
		return nil, err
	}

	usuarioRepository := repository.NewUsuarioRepository(db)
	usuario, err = usuarioRepository.CreateUsuario(usuario)
	if err != nil {
		return nil, err
	}

	return usuario, nil
}
