package models

type Grade struct {
	Id         int `gorm:"primaryKey"`
	Grade      float64
	SubjectID  int
	StudentRUT int
}
