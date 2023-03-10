package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, GreetingMessage{Message: "greeting"})
	})
	e.Logger.Fatal(e.Start(":8080"))
}

type GreetingMessage struct {
	Message string `json:"message"`
}
