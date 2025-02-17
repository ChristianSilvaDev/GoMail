package route

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/controller"
	"github.com/gin-gonic/gin"
)

type RoutesV1 struct {
	PingController *controller.PingController
}

func (r *RoutesV1) Setup(router *gin.Engine) {
	router.GET("api/v1/ping", r.PingController.GetPing)
}
