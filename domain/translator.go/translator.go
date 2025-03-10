package translator

import (
	"nba-players/models/player"
	"nba-players/models/team"
	"nba-players/representations"
	"slices"
)

// converting a rep to team and player models
func Translate(response representations.APIResponse) ([]player.Model, []team.Model) {
	var players []player.Model = make([]player.Model, response.Meta.TotalCount)
	teams := []team.Model{}

	for idx, player := range response.Data {
		players[idx] = ToPlayerModel(player)
		team := players[idx].Team
		if !slices.Contains(teams, team) {
			teams = append(teams, team)
		}
	}

	return players, teams
}

// function to convert a player rep into a domain object
func ToPlayerModel(rep representations.Player) player.Model {
	return player.Model{
		ID:           rep.ID,
		FirstName:    rep.FirstName,
		LastName:     rep.LastName,
		Position:     rep.Position,
		HeightFeet:   rep.HeightFeet,
		HeightInches: rep.HeightInches,
		WeightPounds: rep.WeightPounds,
		Team:         ToTeamModel(rep.Team),
		TeamID:       rep.Team.ID,
	}
}

// function to convert a team rep into a domain object
func ToTeamModel(rep representations.Team) team.Model {
	return team.Model{
		ID:           rep.ID,
		Abbreviation: rep.Abbreviation,
		City:         rep.City,
		Conference:   rep.Conference,
		Division:     rep.Division,
		FullName:     rep.FullName,
		Name:         rep.Name,
	}
}
