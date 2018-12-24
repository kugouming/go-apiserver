package router

import (
	"github.com/gin-gonic/gin"
	"go-apiserver/handler/sd"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine  {
	// Middlewares
	g.Use(gin.Recovery())
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}