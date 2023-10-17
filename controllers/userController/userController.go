package usercontroller

import (
	"belajar-clean/models/dto"
	userservice "belajar-clean/services/userService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Service userservice.UserService
}

func (uc *UserController) GetAllUser(c echo.Context) error  {
	
	data , err := uc.Service.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusExpectationFailed,map[string]interface{}{
			"is_success" : false,
			"message" : "failed get all users",
			"data" : err,
		})
	}

	return c.JSON(http.StatusExpectationFailed,map[string]interface{}{
		"is_success" : true,
		"message" : "success get all users",
		"data" : data,
	})

}

func (us *UserController) CreateUser(c echo.Context) error  {

	input := new(dto.UserCreate)

	c.Bind(&input)

	data , err := us.Service.CreateUser(*input)

	if  err != nil {
		return c.JSON(500, map[string]interface{}{
			"is_success" : false,
			"message" : "failed get all users",
			"data" : err,
		})
	}

	return c.JSON(http.StatusCreated,map[string]interface{}{
		"is_success" : true,
		"message" : "success create user",
		"data" : data,
	})

}