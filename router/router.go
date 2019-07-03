package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"maomitown.com/handler/sd"
	"maomitown.com/handler/user"
	"maomitown.com/router/middleware"
)

// Load 加载路由配置
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	//恢复Api服务器
	g.Use(gin.Recovery())
	//不使用缓存
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	//处理404情况
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "找不到路由")
	})
	//服务器健康检查的路由
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
	}

	return g
}
