package repository

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/dao"
	"github.com/ChristianSilvaDev/GoMail/src/internal/entity"
)

type MailRepository struct {
	*BaseDynamoRepository[entity.Mail]
}

func (r *MailRepository) Create(itemDao dao.MailCreateDAO) (*entity.Mail, error) {
	return r.Create(itemDao)
}
