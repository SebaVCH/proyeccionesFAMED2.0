package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proyeccionesFAMED/database"
	"proyeccionesFAMED/models"
)

func GetAllSubjects(c *gin.Context) {
	var subjects []models.Subject

	if err := database.DB.Find(&subjects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al obtener las asignaturas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subjects})
}
