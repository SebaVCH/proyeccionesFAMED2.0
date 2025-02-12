package models

type Subject struct {
	Id          int `gorm:"primaryKey"`
	NameSubject string
	Credits     int
}
