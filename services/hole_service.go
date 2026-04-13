package services

import (
	"golfscoreid-jng/domains"
	"golfscoreid-jng/logger"
	"golfscoreid-jng/repositories"
)

// HoleService interface defines hole service methods
type HoleService interface {
	GetHoles() ([]domains.Hole, error)
	GetHoleByNumber(number int) (domains.Hole, error)
}

type holeService struct {
	logger logger.Logger
	repo   *repositories.Repository
}

// NewHoleService creates a new hole service instance
func NewHoleService(logger logger.Logger, repo *repositories.Repository) HoleService {
	return &holeService{logger: logger, repo: repo}
}

// GetHoles retrieves all holes ordered by hole number
func (s *holeService) GetHoles() ([]domains.Hole, error) {
	var holes []domains.Hole
	err := s.repo.GetReadDB().Order("hole_number ASC").Find(&holes).Error
	return holes, err
}

// GetHoleByNumber retrieves a hole by its number
func (s *holeService) GetHoleByNumber(number int) (domains.Hole, error) {
	var hole domains.Hole
	err := s.repo.GetReadDB().Where("hole_number = ?", number).First(&hole).Error
	return hole, err
}
