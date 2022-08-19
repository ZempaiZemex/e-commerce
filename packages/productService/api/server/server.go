package server

import "github.com/gin-gonic/gin"

var Engine *gin.Engine = nil

func InitServer() {
	Engine = gin.Default()
}

func Start() {
	if Engine == nil {
		panic("API is not initialized")
	}

	Engine.Run(":4001")
}
