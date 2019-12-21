package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/igson/microservice-go/mvc/service"
	"github.com/igson/microservice-go/mvc/utils"
)

//BuscarUsuario - método de busca do usuário por id.
func BuscarUsuario(resp http.ResponseWriter, req *http.Request) {

	usuarioID, error := strconv.ParseInt(req.URL.Query().Get("usuarioId"), 10, 64)

	if error != nil {

		msgErro := &utils.ErrorMessage{
			Mensagem:   fmt.Sprintf("O campo id deve ser númerico."),
			StatusCode: http.StatusNotFound,
			Codigo:     "erro_conversão_dados",
		}

		jsonMsgErro, _ := json.Marshal(msgErro)
		resp.WriteHeader(msgErro.StatusCode)
		resp.Write(jsonMsgErro)
		return
	}

	usuario, msgErro := service.BuscarUsuario(usuarioID)

	if msgErro != nil {
		jsonMsgErro, _ := json.Marshal(msgErro)
		resp.WriteHeader(msgErro.StatusCode)
		resp.Write([]byte(jsonMsgErro))
		return
	}

	jsonUsuario, _ := json.Marshal(usuario)

	resp.Write(jsonUsuario)

}
