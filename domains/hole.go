package domains

// Hole represents a golf hole configuration (Jatinangor Golf Course)
type Hole struct {
	ID         uint `gorm:"primaryKey;autoIncrement" json:"id"`
	HoleNumber int  `gorm:"column:hole_number;not null;uniqueIndex" json:"hole_number"`
	Par        int  `gorm:"column:par;not null" json:"par"`
	HCP        int  `gorm:"column:hcp;not null" json:"hcp"`
	Black      int  `gorm:"column:black" json:"black"`
	Blue       int  `gorm:"column:blue" json:"blue"`
	White      int  `gorm:"column:white" json:"white"`
	Red        int  `gorm:"column:red" json:"red"`
}

// TableName sets the table name for Hole model
func (Hole) TableName() string {
	return "holes"
}
