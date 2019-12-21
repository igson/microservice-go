package service

import (
	"github.com/igson/microservice-go/mvc/domain"
	"github.com/igson/microservice-go/mvc/repository"
	"github.com/igson/microservice-go/mvc/utils"
)

//BuscarUsuario - método de busca do usuário por id.
func BuscarUsuario(usuarioId int64) (*domain.Usuario, *utils.ErrorMessage) {
	return repository.BuscarUsuario(usuarioId)
}
