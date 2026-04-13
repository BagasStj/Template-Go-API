package controllers

import (
	"net/http"
	"golfscoreid-jng/config"
	"golfscoreid-jng/domains"
	"golfscoreid-jng/logger"
	"golfscoreid-jng/models/errors"
	"golfscoreid-jng/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ScoreController handles score HTTP requests
type ScoreController struct {
	cfg          config.Config
	logger       logger.Logger
	scoreService services.ScoreService
}

// NewScoreController creates a new score controller
func NewScoreController(cfg config.Config, logger logger.Logger, scoreService services.ScoreService) *ScoreController {
	return &ScoreController{cfg: cfg, logger: logger, scoreService: scoreService}
}

// SubmitScore submits a score for a player on a hole (upsert)
// POST /v1/scores
func (sc *ScoreController) SubmitScore(c *gin.Context) {
	var request struct {
		PlayerID   string `json:"player_id" binding:"required"`
		HoleNumber int    `json:"hole_number" binding:"required,min=1,max=18"`
		Strokes    int    `json:"strokes" binding:"required,min=1,max=20"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.SetResponseWithDetails(errors.ErrCodeValidation, err.Error()))
		return
	}

	score := domains.Score{
		PlayerID:   request.PlayerID,
		HoleNumber: request.HoleNumber,
		Strokes:    request.Strokes,
	}

	created, err := sc.scoreService.SubmitScore(score)
	if err != nil {
		sc.logger.Errorf("Failed to submit score: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Score submitted successfully",
		"data":    created,
	})
}

// GetPlayerScores retrieves all scores for a player
// GET /v1/scores?player_id=xxx
func (sc *ScoreController) GetPlayerScores(c *gin.Context) {
	playerID := c.Query("player_id")
	if playerID == "" {
		c.JSON(http.StatusBadRequest, errors.SetResponseWithDetails(errors.ErrCodeBadRequest, "player_id query parameter is required"))
		return
	}

	scores, err := sc.scoreService.GetPlayerScores(playerID)
	if err != nil {
		sc.logger.Errorf("Failed to get player scores: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Scores retrieved successfully",
		"data":    scores,
	})
}

// UpdateScore updates a player's score by score ID
// PUT /v1/scores/:id
func (sc *ScoreController) UpdateScore(c *gin.Context) {
	scoreID := c.Param("id")

	var request struct {
		Strokes int `json:"strokes" binding:"required,min=1,max=20"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.SetResponseWithDetails(errors.ErrCodeValidation, err.Error()))
		return
	}

	updated, err := sc.scoreService.UpdateScore(scoreID, request.Strokes)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, errors.SetResponse(errors.ErrCodeNotFound))
			return
		}
		sc.logger.Errorf("Failed to update score: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Score updated successfully",
		"data":    updated,
	})
}

// DeleteScore deletes a player's score
// DELETE /v1/scores/:id
func (sc *ScoreController) DeleteScore(c *gin.Context) {
	scoreID := c.Param("id")

	if err := sc.scoreService.DeleteScore(scoreID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, errors.SetResponse(errors.ErrCodeNotFound))
			return
		}
		sc.logger.Errorf("Failed to delete score: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Score deleted successfully",
	})
}

// GetLeaderboard retrieves the tournament leaderboard
// GET /v1/leaderboard
func (sc *ScoreController) GetLeaderboard(c *gin.Context) {
	entries, err := sc.scoreService.GetLeaderboard()
	if err != nil {
		sc.logger.Errorf("Failed to get leaderboard: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Leaderboard retrieved successfully",
		"data":    entries,
	})
}
