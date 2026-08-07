package api

import (
	"fmt"

	"github.com/arman-ala/Car_shop/api/routers"
	"github.com/arman-ala/Car_shop/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.SetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	// api group
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			healthGroup := v1.Group("/health")

			routers.HealthRouter(healthGroup)
		}
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
