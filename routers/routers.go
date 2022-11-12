package routers

import (
	"frp-panel/controller"
	docs "frp-panel/docs"
	"frp-panel/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	r.POST("/api/login/account", controller.Login)
	r.POST("/api/login/outLogin", controller.OutLogin)
	r.GET("/api/currentUser", middleware.JWTAuth(), controller.CurrentUser)

	r.GET("/api/frps/config", middleware.JWTAuth(), controller.GetFrpsConfig)
	r.POST("/api/frps/config", middleware.JWTAuth(), controller.SaveFrpsConfig)
	r.GET("/api/frps/info", middleware.JWTAuth(), controller.GetFrpsInfo)
	r.GET("/api/proxy/tcp", middleware.JWTAuth(), controller.GetTcpProxy)
	r.GET("/api/proxy/udp", middleware.JWTAuth(), controller.GetUdpProxy)
	r.GET("/api/proxy/http", middleware.JWTAuth(), controller.GetHttpProxy)
	r.GET("/api/proxy/https", middleware.JWTAuth(), controller.GetHttpsProxy)
	r.GET("/api/proxy/stcp", middleware.JWTAuth(), controller.GetStcpProxy)
	r.GET("/api/proxy/sudp", middleware.JWTAuth(), controller.GetSudpProxy)
	r.GET("/api/proxy/xtcp", middleware.JWTAuth(), controller.GetXtcpProxy)
	r.GET("/api/proxy/tcpmux", middleware.JWTAuth(), controller.GetTcpmuxProxy)
	r.POST("/api/proxy/tcp", middleware.JWTAuth(), controller.AddTcpProxy)
	r.GET("/api/proxy/name/validate", middleware.JWTAuth(), controller.ValidateProxyName)
	r.GET("/api/proxy/port/tcp/validate", middleware.JWTAuth(), controller.ValidateTcpPort)
	r.DELETE("/api/proxy/tcp", middleware.JWTAuth(), controller.DeleteProxy)

	r.GET("/api/frpc", middleware.JWTAuth(), controller.GetFrpc)
	r.GET("/api/frpc/option", middleware.JWTAuth(), controller.GetFrpcOption)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
