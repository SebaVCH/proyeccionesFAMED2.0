package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proyeccionesFAMED/database"
	"proyeccionesFAMED/models"
)

func GetStudentGrades(c *gin.Context) {
	var grades []models.Grade

	rut, exists := c.Get("rut")

	if !exists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error al obtener las calificaciones del estudiante"})
		return
	}

	if err := database.DB.Where("StudentRUT = ?", rut).Find(&grades); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error al obtener las calificaciones del estudiante"})
		return
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"grades": grades})

}

func SimulateGrades(c *gin.Context) {
	rut, exists := c.Get("rut")
	if !exists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error al obtener las calificaciones del estudiante"})
		return
	}
	var grades []models.Grade
	var subjects []models.Subject

	if err := database.DB.Find(&subjects).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error al obtener las asignaturas"})
		return
	}

	if err := database.DB.Where("StudentRUT = ?", rut).Find(&grades).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error al obtener las calificaciones del estudiante"})
		return
	}

	subjectCredits := map[int]int{}
	totalCredits := 0

	for _, subject := range subjects {
		subjectCredits[subject.Id] = subject.Credits
		totalCredits += subjectCredits[subject.Id]
	}

	var weightedSum float64
	var totalAssignedCredits int

	for _, subject := range subjects {
		credit := subjectCredits[subject.Id]
		gradeValue := 0.0
		for _, grade := range grades {
			if grade.SubjectID == subject.Id {
				gradeValue = grade.Grade
				break
			}
		}

		weightedSum += gradeValue * float64(credit)
		totalAssignedCredits += credit
	}

	var finalAverage float64
	if totalAssignedCredits > 0 {
		finalAverage = weightedSum / float64(totalAssignedCredits)
	} else {
		finalAverage = 0
	}
	c.IndentedJSON(http.StatusOK, gin.H{"simulated_average": finalAverage})
}
