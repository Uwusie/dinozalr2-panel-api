package handlers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func WheelGetById(c echo.Context) error {

	pathParam := c.Param("wheelId")

	_, err := strconv.Atoi(pathParam)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid path, please assure that the value is number")
	}

	if pathParam == "" {
		return c.String(http.StatusBadRequest, "Invalid path params")
	}

	path := "Data\\" + pathParam + ".json"
	file, err := os.Open(path)

	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't open file")
	}
	defer file.Close()

	return c.File(path)
}

func WheelsDeleteById(c echo.Context) error {

	pathParam := c.Param("wheelId")

	_, err := strconv.Atoi(pathParam)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid path, please assure that the value is number")
	}

	if pathParam == "" {
		return c.String(http.StatusBadRequest, "Invalid path params")
	}

	file := "Data\\" + pathParam + ".json"
	currentFile, err := os.Open(file)

	if err != nil {
		return c.String(http.StatusBadRequest, "There is no wheel with given id")
	}

	currentFile.Close()

	err = os.Remove(file)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Couldn't delete file")
	}

	return c.String(http.StatusBadRequest, "Wheel deleted properly")
}
