package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-service/config"
	"go-rest-service/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func PutUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	user.Password = string(hashedPassword)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
}
func PostUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	user.Password = string(hashedPassword)
	config.DB.Create(&user)
	c.JSON(201, &user)
}
func GetUserById(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.JSON(200, &user)
}
