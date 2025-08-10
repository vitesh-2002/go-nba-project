package domain

import (
	"nba-players/database"
	"nba-players/models/player"

	"github.com/google/uuid"
)

func InsertPlayer(player *player.Model) error {
	// give players new unique uuids for id
	player.ID = uuid.NewString()
	result := database.DB.Create(player)
	return result.Error
}

func GetPlayerByID(id string) (*player.Model, error) {
	var player player.Model
	result := database.DB.First(&player, id)
	return &player, result.Error
}

func GetPlayersByTeamID(teamID string) ([]player.Model, error) {
	var players []player.Model
	err := database.DB.Where("team_id = ?", teamID).Find(&players).Error
	return players, err
}
