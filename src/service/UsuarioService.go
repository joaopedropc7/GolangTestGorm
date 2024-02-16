package service

import (
	"Routes/src/authentication"
	"Routes/src/banco"
	"Routes/src/models"
	"Routes/src/repository"
	"Routes/src/security"
	"errors"
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

func LoginWithEmail(email string, senha string) (string, error) {
	db, err := banco.Conectar()
	if err != nil {
		return "", err
	}
	usuarioRepository := repository.NewUsuarioRepository(db)

	usuario, err := usuarioRepository.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err = security.VerificarSenha(usuario.Senha, senha); err != nil {
		return "", errors.New("Email ou senha incorreto!")
	}

	token, erro := authentication.CreateToken(usuario.ID)
	if erro != nil {
		return "", erro
	}

	return token, nil

}

func FindUserById(usuarioId uint64) (*models.Usuario, error) {
	db, err := banco.Conectar()
	if err != nil {
		return nil, err
	}

	usuarioRepository := repository.NewUsuarioRepository(db)

	usuario, err := usuarioRepository.GetUserById(usuarioId)
	if err != nil {
		return nil, err
	}

	return &usuario, nil
}
