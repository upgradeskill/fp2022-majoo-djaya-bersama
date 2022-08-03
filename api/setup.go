package api

import (
	"mini-pos/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupResponseGet(c echo.Context, data interface{}, err error) error {
	resp := make(map[string]interface{})

	if err != nil {
		resp["message"] = err.Error()
		return c.JSON(http.StatusNotFound, resp)
	}

	resp["message"] = "Success"
	resp["data"] = data
	return c.JSON(http.StatusOK, resp)
}

func SetupResponsePost(c echo.Context, data interface{}, validate []dto.ValidationMessage, err error) error {
	resp := make(map[string]interface{})

	if validate != nil {
		resp["message"] = "invalid parameters"
		resp["error_validation"] = validate
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		resp["message"] = err.Error()
		return c.JSON(http.StatusNotFound, resp)
	}

	resp["message"] = "Success"
	resp["data"] = data
	return c.JSON(http.StatusOK, resp)
}
