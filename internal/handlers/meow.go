package handlers

import (
	"dinozarl2-panel-api/internal/rabbitmq"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Meow(c echo.Context) error {
	queryParam := c.QueryParam("count")
	pathParam := c.Param("count")
	fmt.Printf("Path param: %v, Query param: %v\n", pathParam, queryParam)

	if queryParam != "" && pathParam != "" {
		return c.String(http.StatusBadRequest, "Invalid params")
	}

	paramValue := queryParam
	if paramValue == "" {
		paramValue = pathParam
	}

	if paramValue == "" {
		return c.String(http.StatusOK, "Meow ")
	}

	meowCount, err := strconv.Atoi(paramValue)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid count")
	}

	return c.String(http.StatusOK, fmt.Sprintf("%s\n", strings.Repeat("Meow ", meowCount)))
}

func ConfigureMeow(c echo.Context) error {
	type MeowConfig struct {
		Count int `json:"count"`
	}

	var config MeowConfig

	if err := c.Bind(&config); err != nil {
		return c.String(400, "Could not parse body")
	}

	rabbitChannel := rabbitmq.GetChannel()

	jsonBody, err := json.Marshal(config)
	if err != nil {
		return c.String(500, "Could not marshal config to JSON")
	}

	err = rabbitChannel.PublishWithContext(c.Request().Context(),
		"",     // exchange
		"meow", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		})

	if err != nil {
		return c.String(500, "Could not publish message")
	}

	return c.String(http.StatusOK, "Message published")
}
