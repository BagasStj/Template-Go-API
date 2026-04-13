package domains

// Score represents a player's score for a specific hole
type Score struct {
	BaseModel
	PlayerID   string `gorm:"column:player_id;type:varchar(36);not null;index" json:"player_id"`
	HoleNumber int    `gorm:"column:hole_number;not null" json:"hole_number"`
	Strokes    int    `gorm:"column:strokes;not null" json:"strokes"`
	Player     Player `gorm:"foreignKey:PlayerID" json:"player,omitempty"`
}

// TableName sets the table name for Score model
func (Score) TableName() string {
	return "scores"
}
