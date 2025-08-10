package service

import (
	"nba-players/client"
	"nba-players/database"
	"nba-players/domain"
	"nba-players/domain/translator.go"
	"nba-players/models/player"
	"nba-players/models/team"
	"nba-players/representations"
)

func Boot() {
	// configure DB
	database.Connect()
	database.DB.AutoMigrate(&team.Model{})
	database.DB.AutoMigrate(&player.Model{})
	// load in all the data we need into our database
	loadData()
}

func FindPlayerByID(id string) (*representations.Player, error) {
	dbPlayer, err := domain.GetPlayerByID(id)
	if err != nil {
		return &representations.Player{}, err
	}
	player := translator.ToPlayerRep(*dbPlayer)
	return &player, nil
}

func FindPlayersByTeamID(id string) ([]representations.Player, error) {
	result := []representations.Player{}
	players, err := domain.GetPlayersByTeamID(id)
	if err != nil {
		return result, err
	}
	for _, player := range players {
		result = append(result, translator.ToPlayerRep(player))
	}

	return result, nil
}

func loadData() error {
	// Send request to get player information
	apiResponse, err := client.GetPlayers()
	if err != nil {
		return err
	}
	// convert representation objects to db models
	players, _ := translator.Translate(apiResponse)
	// insert models into our DB
	for _, player := range players {
		if err := domain.InsertPlayer(&player); err != nil {
			return err
		}
	}

	return nil
}
