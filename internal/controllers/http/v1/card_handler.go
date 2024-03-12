package v1

import (
	"net/http"
	"solidgate-test/config"
	"solidgate-test/internal/controllers/http/requests"
	"solidgate-test/internal/controllers/http/responses"

	"github.com/gin-gonic/gin"
)

type cardHandler struct {
	g   *gin.Engine
	svc CardService
}

func initCardHandler(
	g *gin.Engine,
	svc CardService,
) {
	handler := cardHandler{
		g:   g,
		svc: svc,
	}

	g.POST("/validate", handler.validateCard)
}

func (h cardHandler) validateCard(c *gin.Context) {
	var card requests.Card
	if err := c.BindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest,
			responses.NewValidationErr(config.CodeBadRequest, err.Error()))
	}

	valid, svcCode, err := h.svc.ValidateCard(c.Request.Context(), card)
	if err != nil {
		c.JSON(config.CodeToHttpStatus(svcCode), responses.NewValidationErr(svcCode, err.Error()))
		return
	}
	c.JSON(config.CodeToHttpStatus(svcCode), responses.NewValidationResponse(valid))
}
