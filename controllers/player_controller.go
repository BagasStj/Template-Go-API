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

// PlayerController handles player HTTP requests
type PlayerController struct {
	cfg           config.Config
	logger        logger.Logger
	playerService services.PlayerService
}

// NewPlayerController creates a new player controller
func NewPlayerController(cfg config.Config, logger logger.Logger, playerService services.PlayerService) *PlayerController {
	return &PlayerController{cfg: cfg, logger: logger, playerService: playerService}
}

// RegisterPlayer registers a new player or returns the existing one
// POST /v1/players
func (pc *PlayerController) RegisterPlayer(c *gin.Context) {
	var request struct {
		FullName     string `json:"full_name" binding:"required"`
		BagTagNumber string `json:"bagtag_number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.SetResponseWithDetails(errors.ErrCodeValidation, err.Error()))
		return
	}

	// If bagtag already exists, return the existing player
	existing, err := pc.playerService.GetPlayerByBagTag(request.BagTagNumber)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Player already registered, welcome back!",
			"data":    existing,
		})
		return
	}

	player := domains.Player{
		FullName:     request.FullName,
		BagTagNumber: request.BagTagNumber,
	}

	created, err := pc.playerService.CreatePlayer(player)
	if err != nil {
		pc.logger.Errorf("Failed to create player: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Player registered successfully",
		"data":    created,
	})
}

// GetPlayers retrieves all registered players
// GET /v1/players
func (pc *PlayerController) GetPlayers(c *gin.Context) {
	players, err := pc.playerService.GetPlayers()
	if err != nil {
		pc.logger.Errorf("Failed to get players: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Players retrieved successfully",
		"data":    players,
	})
}

// GetPlayerByID retrieves a player by UUID
// GET /v1/players/:id
func (pc *PlayerController) GetPlayerByID(c *gin.Context) {
	playerID := c.Param("id")

	player, err := pc.playerService.GetPlayerByID(playerID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, errors.SetResponse(errors.ErrCodeNotFound))
			return
		}
		pc.logger.Errorf("Failed to get player: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Player retrieved successfully",
		"data":    player,
	})
}

// GetPlayerByBagTag retrieves a player by bag tag number
// GET /v1/players/bagtag/:bagtag
func (pc *PlayerController) GetPlayerByBagTag(c *gin.Context) {
	bagTag := c.Param("bagtag")

	player, err := pc.playerService.GetPlayerByBagTag(bagTag)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, errors.SetResponse(errors.ErrCodeNotFound))
			return
		}
		pc.logger.Errorf("Failed to get player by bagtag: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Player retrieved successfully",
		"data":    player,
	})
}
