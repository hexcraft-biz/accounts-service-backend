package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hexcraft-biz/accounts-service-backend/config"
	"github.com/hexcraft-biz/accounts-service-backend/features"
)

func main() {
	cfg, err := config.Load()
	MustNot(err)
	cfg.DBOpen(false)

	engine := gin.Default()
	engine.SetTrustedProxies([]string{cfg.Env.TrustProxy})

	// base features
	features.LoadCommon(engine, cfg)
	// auth
	features.LoadAuth(engine, cfg)

	engine.Run(":" + cfg.Env.AppPort)
}

func MustNot(err error) {
	if err != nil {
		panic(err.Error())
	}
}
