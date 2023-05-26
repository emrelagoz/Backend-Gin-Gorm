package main

import (
	"example.com/m/initializers"
	"example.com/m/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	// Kullanıcı tablosu oluşturdum models paketindeki users structı ile
	initializers.DB.AutoMigrate(&models.Users{})

}
