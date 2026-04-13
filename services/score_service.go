package services

import (
	"sort"
	"golfscoreid-jng/domains"
	"golfscoreid-jng/logger"
	"golfscoreid-jng/repositories"
)

// LeaderboardEntry represents a player's leaderboard summary
type LeaderboardEntry struct {
	Player         domains.Player  `json:"player"`
	Scores         []domains.Score `json:"scores"`
	TotalStrokes   int             `json:"total_strokes"`
	TotalPar       int             `json:"total_par"`
	ScoreToPar     int             `json:"score_to_par"`
	HolesCompleted int             `json:"holes_completed"`
}

// ScoreService interface defines score service methods
type ScoreService interface {
	SubmitScore(score domains.Score) (domains.Score, error)
	GetPlayerScores(playerID string) ([]domains.Score, error)
	UpdateScore(id string, strokes int) (domains.Score, error)
	DeleteScore(id string) error
	GetLeaderboard() ([]LeaderboardEntry, error)
}

type scoreService struct {
	logger      logger.Logger
	repo        *repositories.Repository
	holeService HoleService
}

// NewScoreService creates a new score service instance
func NewScoreService(logger logger.Logger, repo *repositories.Repository, holeService HoleService) ScoreService {
	return &scoreService{logger: logger, repo: repo, holeService: holeService}
}

// SubmitScore submits or updates a player's score for a hole
func (s *scoreService) SubmitScore(score domains.Score) (domains.Score, error) {
	var existing domains.Score
	result := s.repo.GetReadDB().Where(
		"player_id = ? AND hole_number = ? AND deleted_at IS NULL",
		score.PlayerID, score.HoleNumber,
	).First(&existing)

	if result.Error == nil {
		// Score already exists — update it
		existing.Strokes = score.Strokes
		err := s.repo.GetWriteDB().Save(&existing).Error
		return existing, err
	}

	err := s.repo.GetWriteDB().Create(&score).Error
	return score, err
}

// GetPlayerScores retrieves all scores for a player
func (s *scoreService) GetPlayerScores(playerID string) ([]domains.Score, error) {
	var scores []domains.Score
	err := s.repo.GetReadDB().
		Where("player_id = ? AND deleted_at IS NULL", playerID).
		Order("hole_number ASC").
		Find(&scores).Error
	return scores, err
}

// UpdateScore updates the strokes for an existing score
func (s *scoreService) UpdateScore(id string, strokes int) (domains.Score, error) {
	var score domains.Score
	err := s.repo.GetWriteDB().Where("id = ? AND deleted_at IS NULL", id).First(&score).Error
	if err != nil {
		return score, err
	}
	score.Strokes = strokes
	err = s.repo.GetWriteDB().Save(&score).Error
	return score, err
}

// DeleteScore soft-deletes a score
func (s *scoreService) DeleteScore(id string) error {
	return s.repo.GetWriteDB().Where("id = ?", id).Delete(&domains.Score{}).Error
}

// GetLeaderboard returns all players ranked by score
func (s *scoreService) GetLeaderboard() ([]LeaderboardEntry, error) {
	var players []domains.Player
	if err := s.repo.GetReadDB().Where("deleted_at IS NULL").Order("created_at ASC").Find(&players).Error; err != nil {
		return nil, err
	}

	holes, err := s.holeService.GetHoles()
	if err != nil {
		return nil, err
	}

	holeParMap := make(map[int]int)
	for _, h := range holes {
		holeParMap[h.HoleNumber] = h.Par
	}

	entries := make([]LeaderboardEntry, 0, len(players))
	for _, player := range players {
		var scores []domains.Score
		s.repo.GetReadDB().Where("player_id = ? AND deleted_at IS NULL", player.ID).Order("hole_number ASC").Find(&scores)

		totalStrokes := 0
		totalPar := 0
		for _, sc := range scores {
			totalStrokes += sc.Strokes
			totalPar += holeParMap[sc.HoleNumber]
		}

		entries = append(entries, LeaderboardEntry{
			Player:         player,
			Scores:         scores,
			TotalStrokes:   totalStrokes,
			TotalPar:       totalPar,
			ScoreToPar:     totalStrokes - totalPar,
			HolesCompleted: len(scores),
		})
	}

	// Sort: most holes completed first, then by lowest strokes
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].HolesCompleted != entries[j].HolesCompleted {
			return entries[i].HolesCompleted > entries[j].HolesCompleted
		}
		return entries[i].TotalStrokes < entries[j].TotalStrokes
	})

	return entries, nil
}
