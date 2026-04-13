package routes

import (
	"golfscoreid-jng/controllers"
	"golfscoreid-jng/services"
)

func (rt *route) initHole() {
	holeService := services.NewHoleService(rt.logger, rt.repo)
	holeController := controllers.NewHoleController(rt.cfg, rt.logger, holeService)

	holes := rt.group.Group("/holes")
	{
		holes.GET("", holeController.GetHoles)
		holes.GET("/:number", holeController.GetHoleByNumber)
	}
}
