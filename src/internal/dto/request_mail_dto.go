package dto

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/dao"
	"github.com/ChristianSilvaDev/GoMail/src/internal/interfaces"
)

type RequestMailDTO struct {
	Destination string `json:"destination"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
}

func (dto *RequestMailDTO) ToDAO() *dao.MailCreateDAO {
	return &dao.MailCreateDAO{
		Destination: dto.Destination,
		Subject:     dto.Subject,
		Body:        dto.Body,
	}
}

var _ interfaces.DTO = (*RequestMailDTO)(nil)
