package routers

import (
	"voting_system/app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	r.Use(cors.New(config))

	v1 := r.Group("/")

	AddRoutesAuth(v1)
	AddRoutesAdministrator(v1)

	v1.Use(middleware.GinJWTMiddlewareHandler().MiddlewareFunc())
	{
		AddRoutesVoter(v1)
	}
	return r
}
