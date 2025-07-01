package router

import "github.com/lin-snow/ech0/internal/di"

// setupEchoRoutes 设置Echo路由
func setupEchoRoutes(appRouterGroup *AppRouterGroup, h *di.Handlers) {
	// Public

	// Auth
	appRouterGroup.AuthRouterGroup.POST("/echo", h.EchoHandler.PostEcho())
	appRouterGroup.AuthRouterGroup.POST("/echo/page", h.EchoHandler.GetEchosByPage())
	appRouterGroup.AuthRouterGroup.DELETE("/echo/:id", h.EchoHandler.DeleteEcho())
	appRouterGroup.AuthRouterGroup.GET("/echo/today", h.EchoHandler.GetTodayEchos())
	appRouterGroup.AuthRouterGroup.PUT("/echo", h.EchoHandler.UpdateEcho())
}
