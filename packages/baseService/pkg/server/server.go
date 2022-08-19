package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine = nil

func InitServer() {
	Engine = gin.Default()
}

func Start(port int) {
	if Engine == nil {
		panic("API is not initialized")
	}

	Engine.Run(fmt.Sprintf(":%d", port))
}
