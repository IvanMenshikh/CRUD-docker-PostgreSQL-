package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Message struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var db *gorm.DB

// Инициализация БД.
func initDB() {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключится к базе данных %v", err)
	}

	db.AutoMigrate(&Message{})

}

func GetHandler(c echo.Context) error {
	var messages []Message

	if err := db.Find(&messages).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not find the message",
		})
	}

	return c.JSON(http.StatusOK, &messages)
}

func PostHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}

	if err := db.Create(&message).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not create the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was successfully create",
	})
}

func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	var updateMessage Message
	if err := c.Bind(&updateMessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Invalid input",
		})
	}

	if err := db.Model(&Message{}).Where("id = ?", id).Update("text", updateMessage.Text).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was update",
	})
}

func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	if err := db.Delete(&Message{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not delete the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was delete",
	})

}

func main() {

	initDB()

	e := echo.New()

	e.GET("/messages", GetHandler)
	e.POST("/messages", PostHandler)
	e.PATCH("/messages/:id", PatchHandler)
	e.DELETE("/messages/:id", DeleteHandler)
	e.Start(":8080")
}
