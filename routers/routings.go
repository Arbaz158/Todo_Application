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

	server.POST("/add-stuff", handlers.AddStuff)
	server.GET("/add-stuff", handlers.Add)

	server.GET("/get-stuff", handlers.GetStuff)
	server.PUT("/update-stuff", handlers.UpdateStuff)
	server.DELETE("/delete-stuff", handlers.DeleteStuff)
	server.Run(":8080")
}
