package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMeow(t *testing.T) {
	tests := []struct {
		name                 string
		path                 string
		queryParam           string
		expected             string
		expectedError        bool
		expectedErrorMessage string
	}{
		{"No params", "/meow", "", "Meow ", false, ""},
		{"Path param", "/meow/5", "", "Meow Meow Meow Meow Meow \n", false, ""},
		{"Query param", "/meow", "count=3", "Meow Meow Meow \n", false, ""},
		{"Both params", "/meow/5", "count=3", "", true, "Invalid params"},
		{"Invalid path param", "/meow/invalid", "", "", true, "Invalid count"},
		{"Invalid query param", "/meow", "count=invalid", "", true, "Invalid count"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fmt.Printf("Test: %v, Path: %v, Query Param: %v\n", test.name, test.path, test.queryParam)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, test.path, nil)
			if test.queryParam != "" {
				req.URL.RawQuery = test.queryParam
			}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if params := extractParams(test.path); len(params) > 0 {
				c.SetPath("/meow/:count")
				c.SetParamNames("count")
				c.SetParamValues(params...)
			} else {
				c.SetPath("/meow")
			}

			err := Meow(c)
			if test.expectedError {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
				assert.Equal(t, test.expectedErrorMessage, rec.Body.String())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.Equal(t, test.expected, rec.Body.String())
			}
		})
	}
}

func extractParams(path string) []string {
	parts := strings.Split(path, "/")
	return parts[2:]
}
