package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type WheelHandler struct {
	DBClient *dynamodb.Client
}

func (h *WheelHandler) WheelGetById(c echo.Context) error {
	type Sector struct {
		Label  string
		Chance float64
	}

	type Wheel struct {
		WheelId   int
		Name      string
		Sectors   []Sector
		ImagePath string
	}

	pathParam := c.Param("wheelId")

	_, err := strconv.Atoi(pathParam)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid path, please assure that the value is number")
	}

	if pathParam == "" {
		return c.String(http.StatusBadRequest, "Invalid path params")
	}

	getItemInput := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"WheelId": &types.AttributeValueMemberN{Value: pathParam},
		},
		TableName: aws.String("FortuneWheelsTable"),
	}

	result, err := h.DBClient.GetItem(context.TODO(), getItemInput)
	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusInternalServerError, "Could not get item from the database")
	}

	var wheel Wheel
	err = attributevalue.UnmarshalMap(result.Item, &wheel)

	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusInternalServerError, "Could not unmarshal item from the database")
	}

	return c.JSON(http.StatusOK, wheel)
}

func (h *WheelHandler) WheelsDeleteById(c echo.Context) error {

	pathParam := c.Param("wheelId")

	_, err := strconv.Atoi(pathParam)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid path, please assure that the value is number")
	}

	if pathParam == "" {
		return c.String(http.StatusBadRequest, "Invalid path params")
	}
	deleteInput := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"WheelId": &types.AttributeValueMemberN{Value: pathParam},
		},
		TableName: aws.String("FortuneWheelsTable"),
	}

	_, err = h.DBClient.DeleteItem(context.TODO(), deleteInput)
	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusInternalServerError, "Could not delete item from the database")

	}

	return c.String(http.StatusOK, "Wheel deleted properly")
}

func (h *WheelHandler) CreateOrUpdateWheel(c echo.Context) error {
	type Sector struct {
		Label  string
		Chance float64
	}

	type Wheel struct {
		WheelId   int
		Name      string
		Sectors   []Sector
		ImagePath string
	}

	var wheel Wheel

	if err := c.Bind(&wheel); err != nil {
		return c.String(400, "Could not parse body")
	}

	if wheel.WheelId == 0 {
		wheel.WheelId = int(uuid.New().ID())
	}
	wheel.ImagePath = "path/to/image.png"

	item, err := attributevalue.MarshalMap(wheel)

	if err != nil {
		fmt.Println(err.Error())
		return c.String(400, "Cannot marshal wheel")
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("FortuneWheelsTable"),
		Item:      item,
	}
	_, err = h.DBClient.PutItem(context.TODO(), input)
	if err != nil {
		fmt.Println(err.Error())
		return c.String(500, "Could not add item to the database")

	}
	return c.String(200, "OK")
}
