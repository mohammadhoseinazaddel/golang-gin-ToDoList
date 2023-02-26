package main

import (
	routes "Todo_List/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.TodoListRoutes(router)
	port := "8000"
	router.Run(":" + port)
}
