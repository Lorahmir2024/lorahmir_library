package status

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/Lorahmir2024/lorahmir_library/pkg/domain/response"
)

func ResponseCase(c echo.Context, status response.Status, data interface{}) error {

	var respData interface{}

	if data != nil {
		respData = data
	} else {
		respData = echo.Map{}
	}

	switch status {
	case response.SearchSuccess, response.CreateSuccess, response.UpdateSuccess, response.DeleteSuccess,
		response.SearchFailed, response.CreateFailed, response.UpdateFailed, response.DeleteFailed:
		return c.JSON(http.StatusOK, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})

	case response.InternalServerError, response.DBQueryError, response.DBRowsAffectedError, response.DBExecutionError, response.DBRowsError, response.DBScanError, response.DBInitError:
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})

	case response.NoContent:
		return c.NoContent(http.StatusNoContent)

	case response.BadRequest:
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})

	case response.Unauthorized:
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})

	case response.Conflict, response.FormatInvalid:
		return c.JSON(http.StatusConflict, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})

	case response.RequestEntityTooLarge:
		return c.JSON(http.StatusRequestEntityTooLarge, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})

	case response.UnsupportedMediaType:
		return c.JSON(http.StatusUnsupportedMediaType, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})

	default:
		status = response.Unknown
		return c.JSON(http.StatusNotImplemented, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})
	}
}
