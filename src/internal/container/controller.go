package container

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/controller"
)

func ProvidePingController() *controller.PingController {
	return &controller.PingController{}
}
