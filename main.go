package main

import (
	"solidgate-test/config"
	"solidgate-test/internal/app"
)

func main() {
	config.FillConfig()
	app.Run(config.Cfg)
}
