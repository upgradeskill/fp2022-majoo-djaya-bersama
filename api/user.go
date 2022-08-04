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
	e.GET("/user", userHandler.UserList)
	e.GET("/user/:ID", userHandler.UserById)
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

func (hand *UserHandler) UserById(c echo.Context) error {
	//claims from jwtToken to UserClaims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.UserClaims)
	resp := make(map[string]interface{})

	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		resp["message"] = "invalid id"
		return c.JSON(http.StatusBadRequest, resp)
	}

	data, validate, err := hand.userUseCase.UserGetById(uint(id), claims)

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
	data.Password = ""
	resp["message"] = "Success Get Data"
	resp["data"] = data
	return c.JSON(http.StatusOK, resp)
}

func (hand *UserHandler) UserList(c echo.Context) error {
	//claims from jwtToken to UserClaims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.UserClaims)
	resp := make(map[string]interface{})

	data, metadata, err := hand.userUseCase.UserList(c, claims)

	if err != nil {
		if err.Error() == "you don't have right" {
			resp["message"] = err.Error()
			return c.JSON(http.StatusUnauthorized, resp)
		}
		resp["message"] = err.Error()
		return c.JSON(http.StatusNotFound, resp)
	}
	if len(data) > 0 {
		resp["metadata"] = metadata
		resp["message"] = "Success Get Data"
		resp["data"] = data
		return c.JSON(http.StatusOK, resp)
	} else {
		resp["metadata"] = metadata
		resp["message"] = "Record Not Found"
		return c.JSON(http.StatusNoContent, resp)
	}
}
