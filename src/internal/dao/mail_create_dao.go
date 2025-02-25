package dao

import "github.com/ChristianSilvaDev/GoMail/src/internal/interfaces"

type MailCreateDAO struct {
	Destination string `json:"destination"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
}

var _ interfaces.DAO = (*MailCreateDAO)(nil)
