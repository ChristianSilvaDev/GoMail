package route

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/controller"
	"github.com/gin-gonic/gin"
)

type RoutesV1 struct {
	PingController *controller.PingController
	MailController *controller.MailController
}

func (r *RoutesV1) Setup(router *gin.Engine) {
	router.GET("api/v1/ping", r.PingController.GetPing)
	router.POST("api/v1/mail", r.MailController.RequestMail)
}
