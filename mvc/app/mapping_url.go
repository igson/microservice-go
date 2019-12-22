package app

import (
	"github.com/igson/microservice-go/mvc/controller"
)

func mapUrls() {
	rota.GET("/usuarios/:userId", controller.BuscarUsuario)
}
