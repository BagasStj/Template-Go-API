package controllers

import (
	"math"
	"net/http"
	"golfscoreid-jng/config"
	"golfscoreid-jng/domains"
	"golfscoreid-jng/logger"
	"golfscoreid-jng/models"
	"golfscoreid-jng/models/errors"
	"golfscoreid-jng/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	cfg         config.Config
	logger      logger.Logger
	userService services.UserService
}

// NewUserController creates a new user controller instance
func NewUserController(cfg config.Config, logger logger.Logger, userService services.UserService) *UserController {
	return &UserController{
		cfg:         cfg,
		logger:      logger,
		userService: userService,
	}
}

// GetUsers retrieves users with pagination
func (uc *UserController) GetUsers(c *gin.Context) {
	var pagination models.PaginationRequest
	if err := c.ShouldBindQuery(&pagination); err != nil {
		uc.logger.Errorf("Failed to bind pagination query: %v", err)
		c.JSON(http.StatusBadRequest, errors.SetResponseWithDetails(errors.ErrCodeBadRequest, err.Error()))
		return
	}

	users, total, err := uc.userService.GetUsers(pagination)
	if err != nil {
		uc.logger.Errorf("Failed to get users: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	// Convert to response format
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, models.ToUserResponse(user))
	}

	// Calculate pagination metadata
	totalPages := int(math.Ceil(float64(total) / float64(pagination.Limit)))
	paginationResponse := models.PaginationResponse{
		CurrentPage:  pagination.Page,
		TotalPages:   totalPages,
		TotalRecords: total,
		HasNext:      pagination.Page < totalPages,
		HasPrev:      pagination.Page > 1,
	}

	response := models.PaginatedResponse{
		Success:    true,
		Message:    "Users retrieved successfully",
		Data:       userResponses,
		Pagination: paginationResponse,
	}

	c.JSON(http.StatusOK, response)
}

// GetUserByID retrieves a user by ID
func (uc *UserController) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, errors.SetResponseWithDetails(errors.ErrCodeBadRequest, "User ID is required"))
		return
	}

	user, err := uc.userService.GetUserByID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, errors.SetResponse(errors.ErrCodeNotFound))
			return
		}
		uc.logger.Errorf("Failed to get user by ID: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	response := models.APIResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    models.ToUserResponse(user),
	}

	c.JSON(http.StatusOK, response)
}

// CreateUser creates a new user
func (uc *UserController) CreateUser(c *gin.Context) {
	var request models.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		uc.logger.Errorf("Failed to bind request: %v", err)
		c.JSON(http.StatusBadRequest, errors.SetResponseWithDetails(errors.ErrCodeValidation, err.Error()))
		return
	}

	// Check if email already exists
	_, err := uc.userService.GetUserByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, errors.SetResponseWithDetails(errors.ErrCodeConflict, "Email already exists"))
		return
	}

	// Create user domain object
	user := domains.User{
		Name:        request.Name,
		Email:       request.Email,
		Username:    request.Username,
		PhoneNumber: request.PhoneNumber,
		IsActive:    true,
	}

	// Create user
	createdUser, err := uc.userService.CreateUser(user)
	if err != nil {
		uc.logger.Errorf("Failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	response := models.APIResponse{
		Success: true,
		Message: "User created successfully",
		Data:    models.ToUserResponse(createdUser),
	}

	c.JSON(http.StatusCreated, response)
}

// UpdateUser updates an existing user
func (uc *UserController) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, errors.SetResponseWithDetails(errors.ErrCodeBadRequest, "User ID is required"))
		return
	}

	var request models.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		uc.logger.Errorf("Failed to bind request: %v", err)
		c.JSON(http.StatusBadRequest, errors.SetResponseWithDetails(errors.ErrCodeValidation, err.Error()))
		return
	}

	// Create update object
	updates := domains.User{
		Name:        request.Name,
		Username:    request.Username,
		PhoneNumber: request.PhoneNumber,
	}

	// Update user
	updatedUser, err := uc.userService.UpdateUser(userID, updates)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, errors.SetResponse(errors.ErrCodeNotFound))
			return
		}
		uc.logger.Errorf("Failed to update user: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	response := models.APIResponse{
		Success: true,
		Message: "User updated successfully",
		Data:    models.ToUserResponse(updatedUser),
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUser soft deletes a user
func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, errors.SetResponseWithDetails(errors.ErrCodeBadRequest, "User ID is required"))
		return
	}

	err := uc.userService.DeleteUser(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, errors.SetResponse(errors.ErrCodeNotFound))
			return
		}
		uc.logger.Errorf("Failed to delete user: %v", err)
		c.JSON(http.StatusInternalServerError, errors.SetResponse(errors.ErrCodeInternalServerError))
		return
	}

	response := models.APIResponse{
		Success: true,
		Message: "User deleted successfully",
	}

	c.JSON(http.StatusOK, response)
}
