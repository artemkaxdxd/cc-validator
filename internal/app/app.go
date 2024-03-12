package app

import (
	"fmt"
	"log"
	"solidgate-test/config"
	v1 "solidgate-test/internal/controllers/http/v1"
	"solidgate-test/internal/service"

	"github.com/gin-gonic/gin"
)

func Run(cfg config.Config) {

	config.InitCorsConfig()

	services := v1.Services{
		CardService: service.NewCardService(),
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	v1.Router(r, cfg, services)

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatalf("run server error: %s \n", err.Error())
	}
}
