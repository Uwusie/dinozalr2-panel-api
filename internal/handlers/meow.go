package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
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
