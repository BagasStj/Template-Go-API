package services

import (
	"golfscoreid-jng/domains"
	"golfscoreid-jng/logger"
	"golfscoreid-jng/repositories"
)

// PlayerService interface defines player service methods
type PlayerService interface {
	GetPlayers() ([]domains.Player, error)
	GetPlayerByID(id string) (domains.Player, error)
	GetPlayerByBagTag(bagTag string) (domains.Player, error)
	CreatePlayer(player domains.Player) (domains.Player, error)
}

type playerService struct {
	logger logger.Logger
	repo   *repositories.Repository
}

// NewPlayerService creates a new player service instance
func NewPlayerService(logger logger.Logger, repo *repositories.Repository) PlayerService {
	return &playerService{logger: logger, repo: repo}
}

// GetPlayers retrieves all players ordered by creation time
func (s *playerService) GetPlayers() ([]domains.Player, error) {
	var players []domains.Player
	err := s.repo.GetReadDB().Where("deleted_at IS NULL").Order("created_at ASC").Find(&players).Error
	return players, err
}

// GetPlayerByID retrieves a player by UUID
func (s *playerService) GetPlayerByID(id string) (domains.Player, error) {
	var player domains.Player
	err := s.repo.GetReadDB().Where("id = ? AND deleted_at IS NULL", id).First(&player).Error
	return player, err
}

// GetPlayerByBagTag retrieves a player by bag tag number
func (s *playerService) GetPlayerByBagTag(bagTag string) (domains.Player, error) {
	var player domains.Player
	err := s.repo.GetReadDB().Where("bagtag_number = ? AND deleted_at IS NULL", bagTag).First(&player).Error
	return player, err
}

// CreatePlayer creates a new player
func (s *playerService) CreatePlayer(player domains.Player) (domains.Player, error) {
	err := s.repo.GetWriteDB().Create(&player).Error
	return player, err
}
