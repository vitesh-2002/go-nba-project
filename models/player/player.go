package player

import "nba-players/models/team"

type Model struct {
	ID           int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	FirstName    string     `gorm:"column:first_name;type:varchar(50);not null" json:"first_name"`
	LastName     string     `gorm:"column:last_name;type:varchar(50);not null" json:"last_name"`
	Position     string     `gorm:"column:position;type:varchar(20);not null" json:"position"`
	HeightFeet   *int       `gorm:"column:height_feet" json:"height_feet"`
	HeightInches *int       `gorm:"column:height_inches" json:"height_inches"`
	WeightPounds *int       `gorm:"column:weight_pounds" json:"weight_pounds"`
	Team         team.Model `gorm:"foreignKey:TeamID" json:"team"`
	TeamID       int        `gorm:"column:team_id;not null" json:"team_id"` // Added for foreign key
}
