package main

import (
	"log"
	"net/http"
	"placeholder-images-generator/internal/server"

	"github.com/labstack/echo/v4"
)

func handler(ctx echo.Context) error {

	response := struct {
		Hello string
		World string `json:"world"`
	}{"hello", "world"}

	return ctx.JSON(http.StatusOK, response)
}

func main() {
	handlers := []server.ServerHandler{{Path: "/", Handler: handler}}

	e, err := server.CreateServer(":3000", handlers)

	if err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":3000"))
}
