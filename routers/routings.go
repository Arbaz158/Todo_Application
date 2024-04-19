package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/todo_application/handlers"
)

func HandlerFunc() {
	server := gin.Default()
	server.POST("/signup", handlers.SignUp)
	server.GET("/home", handlers.Home)

	server.GET("/signup", handlers.Sign)
	server.POST("/", handlers.Login)
	server.GET("/login", handlers.Log)
	server.Run(":8082")
}
