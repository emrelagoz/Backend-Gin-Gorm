package main

import (
	"example.com/m/initializers"
	"example.com/m/middleware"
	"example.com/m/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.Routes()
	r.Use(middleware.CORSMiddleware())
	routes.GetRoutes(r)

	r.Run()
}
