package routes

import (
	"github.com/gin-gonic/gin"
	"proyeccionesFAMED/handlers"
	"proyeccionesFAMED/middleware"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/register", handlers.RegisterStudent)
	r.POST("/register", handlers.LoginStudent)

	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/student", handlers.GetStudentInfo)

	protected.GET("/subjects", handlers.GetAllSubjects)

	protected.GET("/grades", handlers.GetStudentGrades)
	protected.GET("/simulate", handlers.SimulateGrades)

	return r
}
