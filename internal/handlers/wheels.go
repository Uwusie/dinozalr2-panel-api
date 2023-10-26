package handlers

import (
	"encoding/json"
	"fmt"
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

func CreateOrUpdateWheel(c echo.Context) error {
	type Wheel struct {
		Name    string `json:"name"`
		Id      int    `json:"id"`
		Options []struct {
			Description string  `json:"description"`
			Chance      float64 `json:"chance"`
		}
	}

	var wheel Wheel

	if err := c.Bind(&wheel); err != nil {
		return c.String(400, "Cound not parse body")
	}
	filename := fmt.Sprintf("Data/%d.json", wheel.Id)

	data, err := json.MarshalIndent(wheel, "", "  ")
	if err != nil {
		return c.String(500, "Cannot marshal json")
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return c.String(500, "Cannot save file")
	}

	return c.String(200, "OK")
}
