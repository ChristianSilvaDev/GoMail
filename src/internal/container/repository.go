package container

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/repository"
)

func ProvideMailRepository() *repository.MailRepository {
	return &repository.MailRepository{
		DB: ProvideMySQLDatabase(),
	}
}
