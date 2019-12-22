package utils

import "github.com/gin-gonic/gin"

func Response(ctx *gin.Context, status int, body interface{}) {
	switch ctx.GetHeader("Accept") {
	case "application/xml":
		ctx.XML(status, body)
		return
	case "application/json":
		ctx.JSON(status, body)
		return
	}
}

func ResponseError(ctx *gin.Context, erro *ErrorMessage) {
	switch ctx.GetHeader("Accept") {
	case "application/xml":
		ctx.XML(erro.StatusCode, erro)
		return
	case "application/json":
		ctx.JSON(erro.StatusCode, erro)
		return
	}
}
