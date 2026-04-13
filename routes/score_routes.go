package routes

import (
	"golfscoreid-jng/controllers"
	"golfscoreid-jng/services"
)

func (rt *route) initScore() {
	holeService := services.NewHoleService(rt.logger, rt.repo)
	scoreService := services.NewScoreService(rt.logger, rt.repo, holeService)
	scoreController := controllers.NewScoreController(rt.cfg, rt.logger, scoreService)

	// Score endpoints
	scores := rt.group.Group("/scores")
	{
		scores.POST("", scoreController.SubmitScore)
		scores.GET("", scoreController.GetPlayerScores)
		scores.PUT("/:id", scoreController.UpdateScore)
		scores.DELETE("/:id", scoreController.DeleteScore)
	}

	// Leaderboard endpoint
	rt.group.GET("/leaderboard", scoreController.GetLeaderboard)
}
