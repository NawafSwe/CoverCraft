package entity

import "github.com/google/uuid"

type Job struct {
	ID             uuid.UUID `gorm:"primaryKey" json:"id"`
	Company        string    `json:"company"`
	Location       string    `json:"location"`
	Description    string    `json:"description"`
	Qualifications string    `json:"qualifications"`
	Name           string    `json:"name"`
}
