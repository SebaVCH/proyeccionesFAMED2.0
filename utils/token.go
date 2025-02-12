package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"proyeccionesFAMED/config"
	"proyeccionesFAMED/models"
	"time"
)

func GenerateToken(student models.Student) (string, error) {

	claims := jwt.MapClaims{
		"rutStudent": student.Rut,
		"exp":        time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := config.JwtSecret

	return token.SignedString(jwtSecret)

}
