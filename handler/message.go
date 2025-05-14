package handler

import (
	"net/http"
	"strconv"

	"Golang-start/basic/config"
	"Golang-start/basic/domain"

	"github.com/labstack/echo/v4"
)

func GetHandler(c echo.Context) error {
	var messages []domain.Message
	if err := config.DB.Find(&messages).Error; err != nil {
		return c.JSON(http.StatusBadRequest, domain.Response{
			Status:  "Error",
			Message: "Could not find the message",
		})
	}
	return c.JSON(http.StatusOK, &messages)
}

func PostHandler(c echo.Context) error {
	var message domain.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}

	if err := config.DB.Create(&message).Error; err != nil {
		return c.JSON(http.StatusBadRequest, domain.Response{
			Status:  "Error",
			Message: "Could not create the message",
		})
	}

	return c.JSON(http.StatusOK, domain.Response{
		Status:  "Success",
		Message: "Message was successfully create",
	})
}

func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	var updateMessage domain.Message
	if err := c.Bind(&updateMessage); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Response{
			Status:  "Error",
			Message: "Invalid input",
		})
	}

	if err := config.DB.Model(&domain.Message{}).Where("id = ?", id).Update("text", updateMessage.Text).Error; err != nil {
		return c.JSON(http.StatusBadRequest, domain.Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}

	return c.JSON(http.StatusOK, domain.Response{
		Status:  "Success",
		Message: "Message was update",
	})
}

func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	if err := config.DB.Delete(&domain.Message{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, domain.Response{
			Status:  "Error",
			Message: "Could not delete the message",
		})
	}

	return c.JSON(http.StatusOK, domain.Response{
		Status:  "Success",
		Message: "Message was delete",
	})
}
