package services

import (
	"template-go-api/domains"
	"template-go-api/logger"
	"template-go-api/models"
	"template-go-api/repositories"
)

// UserService interface defines user service methods
type UserService interface {
	GetUserByID(id string) (domains.User, error)
	GetUserByEmail(email string) (domains.User, error)
	GetUsers(pagination models.PaginationRequest) ([]domains.User, int64, error)
	CreateUser(user domains.User) (domains.User, error)
	UpdateUser(id string, updates domains.User) (domains.User, error)
	DeleteUser(id string) error
}

// userService implements UserService interface
type userService struct {
	logger logger.Logger
	repo   *repositories.Repository
}

// NewUserService creates a new user service instance
func NewUserService(logger logger.Logger, repo *repositories.Repository) UserService {
	return &userService{
		logger: logger,
		repo:   repo,
	}
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(id string) (domains.User, error) {
	var user domains.User
	err := s.repo.GetReadDB().Where("id = ?", id).First(&user).Error
	return user, err
}

// GetUserByEmail retrieves a user by email
func (s *userService) GetUserByEmail(email string) (domains.User, error) {
	var user domains.User
	err := s.repo.GetReadDB().Where("email = ?", email).First(&user).Error
	return user, err
}

// GetUsers retrieves users with pagination
func (s *userService) GetUsers(pagination models.PaginationRequest) ([]domains.User, int64, error) {
	var users []domains.User
	var total int64

	// Count total records
	err := s.repo.GetReadDB().Model(&domains.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Get paginated users
	offset := (pagination.Page - 1) * pagination.Limit
	err = s.repo.GetReadDB().
		Offset(offset).
		Limit(pagination.Limit).
		Find(&users).Error

	return users, total, err
}

// CreateUser creates a new user
func (s *userService) CreateUser(user domains.User) (domains.User, error) {
	err := s.repo.GetWriteDB().Create(&user).Error
	return user, err
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(id string, updates domains.User) (domains.User, error) {
	var user domains.User

	// First get the existing user
	err := s.repo.GetWriteDB().Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	// Update the user
	err = s.repo.GetWriteDB().Model(&user).Updates(updates).Error
	return user, err
}

// DeleteUser soft deletes a user
func (s *userService) DeleteUser(id string) error {
	return s.repo.GetWriteDB().Where("id = ?", id).Delete(&domains.User{}).Error
}
