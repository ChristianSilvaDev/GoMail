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
	mail := entity.NewMail(params.Destination, params.Subject, params.Body)
	err := u.MailRepository.Create(mail)
	if err != nil {
		return nil, err
	}

	return mail, nil
}
