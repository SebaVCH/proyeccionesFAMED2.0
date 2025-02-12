package models

type Student struct {
	Rut      string `gorm:"primaryKey"`
	Password string
	Name     string
	LastName string
}
