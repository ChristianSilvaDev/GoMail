package usecase

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/dto"
	"github.com/ChristianSilvaDev/GoMail/src/internal/entity"
	"github.com/ChristianSilvaDev/GoMail/src/internal/repository"
)

type RequestMail struct {
	MailRepository *repository.MailRepository
}

func (u *RequestMail) Execute(params dto.RequestMailDTO) (*entity.Mail, error) {
	return u.MailRepository.Create(*params.ToDAO())
}
