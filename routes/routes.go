package routes

import (
	"example.com/m/controllers"
	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {

	r.POST("/users", controllers.UsersCreate)
	r.GET("/users", controllers.UsersIndex)
	r.GET("/users/:id", controllers.UsersShow)
	r.PUT("/users/:id", controllers.UsersUpdate)
	r.DELETE("/users/:id", controllers.UsersDelete)
}
