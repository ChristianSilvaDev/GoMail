package entity

import (
	"github.com/google/uuid"
	"time"
)

type MailStatus string

const (
	pending MailStatus = "Pending"
	sent    MailStatus = "Sent"
	failed  MailStatus = "Failed"
)


type Mail struct {
	ID          uuid.UUID  `gorm:"type:varchar(36);primary_key" json:"id"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
	Destination string     `gorm:"not null; not empty" json:"destination"`
	Subject     string     `gorm:"not null; not empty" json:"subject"`
	Body        string     `gorm:"not null; not empty" json:"body"`
	Status      MailStatus `gorm:"not null;index" json:"status"`
}

func NewMail(Destination, Subject, Body string) *Mail {
	return &Mail{
		ID:          uuid.New(),
		Destination: Destination,
		Subject:     Subject,
		Body:        Body,
		Status:      pending,
	}
}
