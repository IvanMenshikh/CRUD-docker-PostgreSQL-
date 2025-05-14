package main

import (
	"Golang-start/basic/config"
	"Golang-start/basic/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()

	e := echo.New()
	e.GET("/messages", handler.GetHandler)
	e.POST("/messages", handler.PostHandler)
	e.PATCH("/messages/:id", handler.PatchHandler)
	e.DELETE("/messages/:id", handler.DeleteHandler)
	e.Start(":8080")
}
