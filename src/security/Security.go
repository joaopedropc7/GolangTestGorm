package security

import "golang.org/x/crypto/bcrypt"

// Hash HASH RECEBE UMA STRING E COLOCA UM HASH NELA
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha COMPARA UMA SENHA E UM HASH E RETORNA SE ELAS SAO IGUAIS
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
