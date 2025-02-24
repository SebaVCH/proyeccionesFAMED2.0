package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"proyeccionesFAMED/database"
	"proyeccionesFAMED/models"
)

func GetAllSubjects(c *gin.Context) {
	var subjects []models.Subject

	cacheKey := "all_subjects"
	cachedSubjects, err := database.RedisClient.Get(database.Ctx, cacheKey).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cachedSubjects), &subjects); err == nil {
			c.IndentedJSON(http.StatusOK, gin.H{"data": subjects})
			return
		}
	}

	if err := database.DB.Find(&subjects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al obtener las asignaturas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subjects})
}
