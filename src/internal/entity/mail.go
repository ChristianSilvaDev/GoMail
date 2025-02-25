package entity

import "github.com/ChristianSilvaDev/GoMail/src/internal/interfaces"

type Mail struct {
	*BaseEntity
	Destination string `json:"destination"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	Status      string `json:"status"`
}

var _ interfaces.Entity = (*Mail)(nil)
