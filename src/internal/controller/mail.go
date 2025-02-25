package controller

import (
	"net/http"

	"github.com/ChristianSilvaDev/GoMail/src/internal/dto"
	"github.com/ChristianSilvaDev/GoMail/src/internal/usecase"
	"github.com/gin-gonic/gin"
)

type MailController struct {
	RequestMailUseCase *usecase.RequestMail
}

func (c *MailController) RequestMail(ctx *gin.Context) {
	var requestDTO dto.RequestMailDTO

	if err := ctx.ShouldBindJSON(&requestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mail, err := c.RequestMailUseCase.Execute(requestDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": mail.ID})
}
