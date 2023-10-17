package routes

import (
	usercontroller "belajar-clean/controllers/userController"
	userrepository "belajar-clean/repositories/userRepository"
	userservice "belajar-clean/services/userService"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteService(db *gorm.DB) *echo.Echo {

	//route repository
	userRouteRepository := userrepository.NewUserRepository(db)

	//route service
	userRouteService := userservice.NewUserService(userRouteRepository)

	//route controller
	userRouteController := usercontroller.UserController{
		Service: userRouteService,
	}

	app := echo.New()

	app.GET("/users",userRouteController.GetAllUser)
	app.POST("/user",userRouteController.CreateUser)

	return app

}