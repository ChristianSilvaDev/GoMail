package repository

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/entity"
	"github.com/ChristianSilvaDev/GoMail/src/internal/infra"
)

type MailRepository struct {
	DB *infra.MySQLDatabase
}

func (r *MailRepository) Create(mail *entity.Mail) error {
	conn, err := r.DB.GetConnection()
	if err != nil {
		return err
	}

	conn.Create(mail)

	return nil
}
