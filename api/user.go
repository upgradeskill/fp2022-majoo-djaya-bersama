package api

import (
	"mini-pos/dto"
	"mini-pos/repository"
	"mini-pos/security"
	"mini-pos/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler() *UserHandler {
	userRepo := repository.InitUserRepository()
	authorizeRepo := repository.InitAuthorizeRepository()
	jwtService := security.JWTAuthService()
	userUseCase := usecase.InitUserUseCase(userRepo, authorizeRepo, jwtService)
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func Authorization(e *echo.Group) {

	userHandler := NewUserHandler()

	e.POST("/login", userHandler.Login)
	e.POST("/register", userHandler.Register)
}

func UserApi(e *echo.Group) {
	userHandler := NewUserHandler()
	e.POST("/user", userHandler.UserInsert)
	e.PUT("/user", userHandler.UserUpdate)
	e.DELETE("/user/:ID", userHandler.UserDelete)
}

func (hand *UserHandler) Login(c echo.Context) error {
	resp := make(map[string]interface{})
	data, validate, err := hand.userUseCase.Login(c)

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

func (hand *UserHandler) Register(c echo.Context) error {
	resp := make(map[string]interface{})
	data, validate, err := hand.userUseCase.Register(c)

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

func (hand *UserHandler) UserInsert(c echo.Context) error {
	//claims from jwtToken to UserClaims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.UserClaims)

	resp := make(map[string]interface{})
	data, validate, err := hand.userUseCase.UserInsert(c, claims)

	if validate != nil {
		resp["message"] = "invalid parameters"
		resp["error_validation"] = validate
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		if err.Error() == "you don't have right" {
			resp["message"] = err.Error()
			return c.JSON(http.StatusUnauthorized, resp)
		}
		resp["message"] = err.Error()
		return c.JSON(http.StatusNotFound, resp)
	}

	resp["message"] = "Success"
	resp["data"] = data
	return c.JSON(http.StatusOK, resp)
}

func (hand *UserHandler) UserUpdate(c echo.Context) error {
	//claims from jwtToken to UserClaims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.UserClaims)

	resp := make(map[string]interface{})
	data, validate, err := hand.userUseCase.UserUpdate(c, claims)

	if validate != nil {
		resp["message"] = "invalid parameters"
		resp["error_validation"] = validate
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		if err.Error() == "you don't have right" {
			resp["message"] = err.Error()
			return c.JSON(http.StatusUnauthorized, resp)
		}
		resp["message"] = err.Error()
		return c.JSON(http.StatusNotFound, resp)
	}

	resp["message"] = "Update Success"
	resp["data"] = data
	return c.JSON(http.StatusOK, resp)
}

func (hand *UserHandler) UserDelete(c echo.Context) error {
	//claims from jwtToken to UserClaims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.UserClaims)
	resp := make(map[string]interface{})

	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		resp["message"] = "invalid id"
		return c.JSON(http.StatusBadRequest, resp)
	}

	data, validate, err := hand.userUseCase.UserDelete(uint(id), claims)

	if validate != nil {
		resp["message"] = "invalid parameters"
		resp["error_validation"] = validate
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		if err.Error() == "you don't have right" {
			resp["message"] = err.Error()
			return c.JSON(http.StatusUnauthorized, resp)
		}
		resp["message"] = err.Error()
		return c.JSON(http.StatusNotFound, resp)
	}

	resp["message"] = "Success Delete Data"
	resp["data"] = data
	return c.JSON(http.StatusOK, resp)
}
