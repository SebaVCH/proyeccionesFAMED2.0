package routes

import (
	"github.com/gin-gonic/gin"
	"proyeccionesFAMED/handlers"
	"proyeccionesFAMED/middleware"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/register", handlers.RegisterStudent)
	r.POST("/login", handlers.LoginStudent)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/student", handlers.GetStudentInfo)

	protected.GET("/subjects", handlers.GetAllSubjects)

	protected.GET("/grades", handlers.GetStudentGrades)
	protected.GET("/simulate", handlers.SimulateGrades)

	return r
}
