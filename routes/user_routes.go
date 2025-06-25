package routes

import (
	"template-go-api/controllers"
	"template-go-api/services"
)

// initUser initializes user routes
func (rt *route) initUser() {
	userService := services.NewUserService(rt.logger, rt.repo)
	userController := controllers.NewUserController(rt.cfg, rt.logger, userService)

	userGroup := rt.group.Group("/users")
	{
		userGroup.GET("", userController.GetUsers)
		userGroup.GET("/:id", userController.GetUserByID)
		userGroup.POST("", userController.CreateUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}
}
