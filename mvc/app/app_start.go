package app

import (
	"github.com/gin-gonic/gin"
)

var (
	rota *gin.Engine
)

func init() {
	rota = gin.Default()
}

//StartApp inicia a execução da aplicação
func StartApp() {

	mapUrls()

	if error := rota.Run(":8080"); error != nil {
		panic(error)
	}

}
