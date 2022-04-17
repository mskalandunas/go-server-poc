package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleGet(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", handleGet)
	e.Logger.Fatal(e.Start(":3000"))
}
