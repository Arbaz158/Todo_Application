package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/todo_application/handlers"
)

func HandlerFunc() {
	server := gin.Default()

	server.POST("/signup", handlers.SignUp)
	server.GET("/signup", handlers.Sign)

	server.POST("/login", handlers.Login)
	server.GET("/login", handlers.Log)

	server.GET("/home", handlers.Home)
	server.POST("/add-data", handlers.UpdateEmployee)
	server.Run(":8082")
}
