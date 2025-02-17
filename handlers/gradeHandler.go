package handlers

import (
	"github.com/gin-gonic/gin"
	"math"
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

	if err := database.DB.Where("student_rut = ?", rut).Find(&grades).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Esta persona no tiene calificaciones"})
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

	if err := database.DB.Where("student_rut = ?", rut).Find(&grades).Error; err != nil {
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
		gradeValue := 0.0000
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
		finalAverage = math.Round((weightedSum/float64(totalAssignedCredits))*10000) / 10000
	} else {
		finalAverage = 0
	}
	formattedAverage := math.Round(finalAverage*100) / 100
	c.IndentedJSON(http.StatusOK, gin.H{"simulated_average": formattedAverage})
}
