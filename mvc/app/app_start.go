package app

import (
	"net/http"

	"github.com/igson/microservice-go/mvc/controller"
)

//StartApp inicia a execução da aplicação
func StartApp() {

	http.HandleFunc("/usuarios", controller.BuscarUsuario)

	if error := http.ListenAndServe(":8080", nil); error != nil {
		panic(error)
	}

}
