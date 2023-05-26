package controllers

import (

	"example.com/m/initializers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func UsersCreate(c *gin.Context) {

	// get data off req body
	var body struct {
		FirstName string
		LastName  string
		Email     string
		Phone     string
	}

	c.Bind(&body)

	// user oluşturma
	user := models.Users{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Phone: body.Phone}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// returnleme
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersIndex(c *gin.Context) {

	var users []models.Users
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func UsersShow(c *gin.Context) {
	// urlden id alma
	id := c.Param("id")
	

	// kullanıcıyı idye göre alma
	var user models.Users
	initializers.DB.Find(&user, id)

	// return
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		FirstName string
		LastName  string
		Email     string
		Phone     string
	}

	c.Bind(&body)

	var user models.Users
	initializers.DB.First(&user, id)

	// güncelleme
	initializers.DB.Model(&user).Updates(models.Users{
		FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Phone: body.Phone})

	// yanıt return
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersDelete(c *gin.Context) {

	id := c.Param("id")

	var user models.Users

	// kullanıcı silme
	initializers.DB.Delete(&user, id)

	c.Status(200)
}
