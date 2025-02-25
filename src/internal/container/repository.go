package container

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/entity"
	"github.com/ChristianSilvaDev/GoMail/src/internal/interfaces"
	"github.com/ChristianSilvaDev/GoMail/src/internal/repository"
)

func ProvideBaseDynamoRepository[T interfaces.Entity](
	newEntityFunc func() *T,
	tableName string,
) *repository.BaseDynamoRepository[T] {
	return &repository.BaseDynamoRepository[T]{
		TableName: tableName,
		NewEntity: newEntityFunc,
	}
}

func ProvideMailRepository() *repository.MailRepository {
	newEntityFunc := func() *entity.Mail {
		return &entity.Mail{}
	}

	return &repository.MailRepository{
		BaseDynamoRepository: ProvideBaseDynamoRepository(newEntityFunc, "mails"),
	}
}
