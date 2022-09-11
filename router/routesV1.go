package router

import (
	v1 "backChannel/api/v1"
	"github.com/labstack/echo/v4"
)

func V1Routes(group *echo.Group) {
	group.POST("/auth", v1.BackChannelController().Create)
	group.GET("/auth", v1.BackChannelController().Get)
}
