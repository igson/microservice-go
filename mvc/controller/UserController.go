package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/igson/microservice-go/mvc/service"
	"github.com/igson/microservice-go/mvc/utils"
)

//BuscarUsuario - método de busca do usuário por id.
func BuscarUsuario(ctx *gin.Context) {

	usuarioID, error := strconv.ParseInt(ctx.Param("userId"), 10, 64)

	if error != nil {
		msgErro := &utils.ErrorMessage{
			Mensagem:   fmt.Sprintf("O campo id deve ser númerico."),
			StatusCode: http.StatusNotFound,
			Codigo:     "erro_conversão_dados",
		}
		utils.ResponseError(ctx, msgErro)
		return
	}

	usuarioEncontrado, msgErro := service.BuscarUsuario(usuarioID)

	if msgErro != nil {
		utils.ResponseError(ctx, msgErro)
		return
	}

	utils.Response(ctx, http.StatusOK, usuarioEncontrado)

}
