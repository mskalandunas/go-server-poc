package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Notification struct {
	Body string `json:"body"`
	Href string `json:"name"`
	Icon string `json:"icon"`
}

func handleGetNotifications(c echo.Context) error {
	return c.JSON(http.StatusOK, []Notification{{
		Body: "Text",
		Href: "https://google.com",
		Icon: "ICON_IDENTIFIER",
	}})
}

func main() {
	e := echo.New()
	e.GET("notifications", handleGetNotifications)
	e.Logger.Fatal(e.Start(":3000"))
}
