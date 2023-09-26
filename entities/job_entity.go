package entities

import "github.com/google/uuid"

type Job struct {
	ID             uuid.UUID `gorm:"primaryKey" json:"id"`
	Company        string
	Location       string
	Description    string
	Qualifications string
	Name           string
}
