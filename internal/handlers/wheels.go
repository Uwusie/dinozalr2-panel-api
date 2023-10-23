package handlers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Wheels(c echo.Context) error {

	pathParam := c.Param("wheelsId")

	_, err := strconv.Atoi(pathParam)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid path, please assure that the value is number")
	}

	if pathParam == "" {
		return c.String(http.StatusBadRequest, "Invalid path params")
	}

	file := "Data\\" + pathParam + ".json"
	_, err = os.Open(file)

	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't open file")
	}

	return c.File(file)
}
