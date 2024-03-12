package v1

import (
	"net/http"
	"solidgate-test/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(
	g *gin.Engine,
	cfg config.Config,
	services Services,
) {
	g.Use(
		gin.Logger(),
		gin.Recovery(),
		cors.New(config.CorsConfig),
	)

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	initCardHandler(g, services.CardService)
}
