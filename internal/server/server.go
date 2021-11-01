package server

import (
	"github.com/labstack/echo/v4"
)

type Server interface {
	CreateServer(url string, handlers []ServerHandler)
}

type ServerHandler struct {
	Path    string
	Handler echo.HandlerFunc
}

func CreateServer(url string, handlers []ServerHandler) (*echo.Echo, error) {

	e := echo.New()

	for _, h := range handlers {
		e.GET(h.Path, h.Handler)
	}

	return e, nil
}
