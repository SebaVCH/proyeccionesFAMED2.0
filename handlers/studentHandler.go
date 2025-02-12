package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"proyeccionesFAMED/database"
	"proyeccionesFAMED/models"
	"proyeccionesFAMED/utils"
)

func RegisterStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "Error al registrar el usuario"})
		return
	}

	student.Password = utils.HashingPassword(student.Password)
	database.DB.Create(&student)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Usuario registrado correctamente"})
}
func LoginStudent(c *gin.Context) {
	var input struct {
		Rut      string
		Password string
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "Formato o caracteres invalidos"})
		return
	}

	var student models.Student

	if err := database.DB.Where("rut = ?", input.Rut).First(&student).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "Error al iniciar sesion"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(input.Password)); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "Error al iniciar sesion"})
		return
	}

	token, err := utils.GenerateToken(student)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "Error al iniciar sesion"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})

}
func GetStudentInfo(c *gin.Context) {
	var student models.Student

	c.IndentedJSON(http.StatusBadRequest, &student)

}
