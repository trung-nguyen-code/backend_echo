package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	server := echo.New()
	server.GET("/", hello)
	server.logger.Fatal(server.Start(":8888"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
