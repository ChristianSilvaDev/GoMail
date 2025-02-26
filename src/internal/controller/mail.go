package controller

import (
	"github.com/ChristianSilvaDev/GoMail/src/internal/dto"
	"github.com/ChristianSilvaDev/GoMail/src/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MailController struct {
	RequestMailUseCase *usecase.RequestMail
}

func (c *MailController) RequestMail(ctx *gin.Context) {
	var inputDTO dto.RequestMailDTO

	if err := ctx.ShouldBindJSON(&inputDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mail, err := c.RequestMailUseCase.Execute(inputDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, dto.RequestMailResponseDTO{
		ID:          mail.ID,
		Destination: mail.Destination,
		Subject:     mail.Subject,
	})
}
