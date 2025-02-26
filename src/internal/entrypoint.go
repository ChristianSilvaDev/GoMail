package internal

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/container"
	"github.com/ChristianSilvaDev/GoMail/src/internal/entity"
)

func Setup() {
	db := container.ProvideMySQLDatabase()
	db.Init(
		&entity.Mail{},
	)
}
