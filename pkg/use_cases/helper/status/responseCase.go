package status

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ResponseCase(c echo.Context, data interface{}, error string) error {

	var respData interface{}

	if data != nil {
		respData = data
	} else {
		respData = echo.Map{}
	}

	var errors []string
	if error != "" {
		errors = append(errors, error)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data":   respData,
		"errors": errors,
	})
}
