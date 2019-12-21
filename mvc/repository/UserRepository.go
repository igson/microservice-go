package repository

import (
	"fmt"
	"net/http"

	"github.com/igson/microservice-go/mvc/domain"
	"github.com/igson/microservice-go/mvc/utils"
)

var (
	usuarios = map[int64]*domain.Usuario{
		123: {ID: 1, Nome: "Igson", Sobrenome: "Mendes", Email: "igson@gmail.com"},
	}
)

//BuscarUsuario - método de busca do usuário por id.
func BuscarUsuario(usuarioID int64) (*domain.Usuario, *utils.ErrorMessage) {
	if usuario := usuarios[usuarioID]; usuario != nil {
		return usuario, nil
	}
	return nil, &utils.ErrorMessage{
		Mensagem:   fmt.Sprintf("Usuário de código %d não encontrado.", usuarioID),
		StatusCode: http.StatusNotFound,
		Codigo:     "id_nao_encontrado",
	}
}
