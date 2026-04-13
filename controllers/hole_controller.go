package controllers

import (
	"net/http"
	"strconv"
	"golfscoreid-jng/config"
	"golfscoreid-jng/logger"
	"golfscoreid-jng/models/errors"
	"golfscoreid-jng/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HoleController handles hole HTTP requests
type HoleController struct {
	cfg         config.Config
	logger      logger.Logger
	holeService services.HoleService
}

// NewHoleController creates a new hole controller
func NewHoleController(cfg config.Config, logger logger.Logger, holeService services.HoleService) *HoleController {
	return &HoleController{cfg: cfg, logger: logger, holeService: holeService}
}

// GetHoles retrieves all holes for Jatinangor Golf Course
// GET /v1/holes
func (hc *HoleController) GetHoles(c *gin.Context) {
	holes, err := hc.holeService.GetHoles()
	if err != nil {
		hc.logger.Errorf("Failed to get holes: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Holes retrieved successfully",
		"data":    holes,
	})
}

// GetHoleByNumber retrieves a specific hole by number
// GET /v1/holes/:number
func (hc *HoleController) GetHoleByNumber(c *gin.Context) {
	numberStr := c.Param("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil || number < 1 || number > 18 {
		c.JSON(http.StatusBadRequest, errors.SetResponseWithDetails(errors.ErrCodeBadRequest, "Invalid hole number (must be 1-18)"))
		return
	}

	hole, err := hc.holeService.GetHoleByNumber(number)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, errors.SetResponse(errors.ErrCodeNotFound))
			return
		}
		hc.logger.Errorf("Failed to get hole: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Hole retrieved successfully",
		"data":    hole,
	})
}
