package container

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/controller"
)

func ProvidePingController() *controller.PingController {
	return &controller.PingController{}
}

func ProvideMailController() *controller.MailController {
	return &controller.MailController{
		RequestMailUseCase: ProvideRequestMailUseCase(),
	}
}
