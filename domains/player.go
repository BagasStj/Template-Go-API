package domains

// Player represents a golf player in the system
type Player struct {
	BaseModel
	FullName     string `gorm:"column:full_name;type:varchar(255);not null" json:"full_name"`
	BagTagNumber string `gorm:"column:bagtag_number;type:varchar(50);not null;uniqueIndex" json:"bagtag_number"`
}

// TableName sets the table name for Player model
func (Player) TableName() string {
	return "players"
}
