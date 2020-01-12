package controllers

import (
	"github.com/mnemonic111/poc_api_rest_go/dtos"
)

// UserController  Estructura que contiene el controlador de usuarios y sus dependencias.
type UserController struct {
	Usuario dtos.User
}

// SetNewUser Devuelve un numero aleatorio.
func (s *UserController) SetNewUser(nombre string, apellidos string, email string, password string) (dtos.User) {
	s.Usuario.Nombre = nombre
	s.Usuario.Apellidos = apellidos
	s.Usuario.Email = email
	s.Usuario.Password = password
	return s.Usuario;
}
