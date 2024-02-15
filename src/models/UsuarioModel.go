package models

import (
	"Routes/src/security"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

// Usuario REPRESENTA UM Usuario UTILIZANDO A REDE SOCIAL
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

type UsuarioRequestVO struct {
	Nome  string `json:"nome,omitempty"`
	Nick  string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
}

// Prepara vai chamar os metedos para validar e formatar o Usuario recebido
func (Usuario *Usuario) Preparar(etapa string) error {
	if erro := Usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := Usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (Usuario *Usuario) validar(etapa string) error {
	if Usuario.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}
	if Usuario.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}
	if Usuario.Email == "" {
		return errors.New("o email é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(Usuario.Email); erro != nil {
		return errors.New("o email inserido e invalido")
	}

	if etapa == "cadastro" && Usuario.Senha == "" {
		return errors.New("a senha é obrigatório e não pode estar em branco")
	}

	return nil
}

func (Usuario *Usuario) formatar(etapa string) error {
	Usuario.Nome = strings.TrimSpace(Usuario.Nome)
	Usuario.Nick = strings.TrimSpace(Usuario.Nick)
	Usuario.Email = strings.TrimSpace(Usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := security.Hash(Usuario.Senha)
		if erro != nil {
			return erro
		}

		Usuario.Senha = string(senhaComHash)
	}
	return nil
}
