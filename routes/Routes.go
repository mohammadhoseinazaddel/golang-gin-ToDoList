package routes

import (
	controller "Todo_List/controllers"

	"github.com/gin-gonic/gin"
)

// TodoListRoutes function
func TodoListRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/todo_list", controller.TodoListPost())
	incomingRoutes.GET("/users/todo_list", controller.TodoListGet())
	incomingRoutes.PUT("/users/update", controller.TodoListUpdate())
	// incomingRoutes.DELETE("/users/todo_list", controller.TodoListDelete())
}
