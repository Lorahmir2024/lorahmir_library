package status

import (
	"github.com/Lorahmir2024/lorahmir_library/pkg/domain/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ResponseCase(c echo.Context, data interface{}, status response.Status) error {

	var respData interface{}

	if data != nil {
		respData = data
	} else {
		respData = echo.Map{}
	}

	var errors []string
	if status != response.OK {
		errors = append(errors, status.String())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data":   respData,
		"errors": errors,
	})
}
