package routes

import (
	"golfscoreid-jng/controllers"
	"golfscoreid-jng/services"
)

func (rt *route) initPlayer() {
	playerService := services.NewPlayerService(rt.logger, rt.repo)
	playerController := controllers.NewPlayerController(rt.cfg, rt.logger, playerService)

	players := rt.group.Group("/players")
	{
		players.POST("", playerController.RegisterPlayer)
		players.GET("", playerController.GetPlayers)
		players.GET("/:id", playerController.GetPlayerByID)
		players.GET("/bagtag/:bagtag", playerController.GetPlayerByBagTag)
	}
}
