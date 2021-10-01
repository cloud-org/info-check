package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func addMiddleware(e *echo.Echo) {
	// 增加 cors 中间件
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func addApi(e *echo.Echo, data []PersonInfo) {
	info := NewInfo(data)
	e.GET("/api", info.HandlerInfo)
}

//CreateEngine echo
func CreateEngine(data []PersonInfo) (*echo.Echo, error) {
	e := echo.New()

	addMiddleware(e)
	addApi(e, data)

	return e, nil
}
