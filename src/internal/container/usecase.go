package container

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/usecase"
)

func ProvideRequestMailUseCase() *usecase.RequestMail {
	return &usecase.RequestMail{
		MailRepository: ProvideMailRepository(),
	}
}

