package translator

import (
	"nba-players/models/player"
	"nba-players/models/team"
	"nba-players/representations"
	"strconv"
)

// converting a rep to team and player models
func Translate(response representations.APIResponse) ([]player.Model, []team.Model) {
	var players []player.Model = make([]player.Model, response.Meta.TotalCount)
	teams := []team.Model{}

	for _, player := range response.Data {
		players = append(players, ToPlayerModel(player))
		// team := ToTeamModel(player.Team)
		// if !slices.Contains(teams, team) {
		// 	teams = append(teams, team)
		// }
	}

	return players, teams
}

// function to convert a player rep into a domain object
func ToPlayerModel(rep representations.Player) player.Model {
	return player.Model{
		ID:           intToString(rep.ID),
		FirstName:    rep.FirstName,
		LastName:     rep.LastName,
		Position:     rep.Position,
		HeightFeet:   rep.HeightFeet,
		HeightInches: rep.HeightInches,
		WeightPounds: rep.WeightPounds,
		Team:         ToTeamModel(rep.Team),
		TeamID:       intToString(rep.Team.ID),
	}
}

// function to convert a team rep into a domain object
func ToTeamModel(rep representations.Team) team.Model {
	return team.Model{
		ID:           intToString(rep.ID),
		Abbreviation: rep.Abbreviation,
		City:         rep.City,
		Conference:   rep.Conference,
		Division:     rep.Division,
		FullName:     rep.FullName,
		Name:         rep.Name,
	}
}

func ToPlayerRep(model player.Model) representations.Player {
	return representations.Player{
		ID:           stringToInt(model.ID),
		FirstName:    model.FirstName,
		LastName:     model.LastName,
		Position:     model.Position,
		HeightFeet:   model.HeightFeet,
		HeightInches: model.HeightInches,
		WeightPounds: model.WeightPounds,
		Team:         ToTeamRep(model.Team),
	}
}

func ToTeamRep(model team.Model) representations.Team {
	return representations.Team{
		ID:           stringToInt(model.ID),
		Abbreviation: model.Abbreviation,
		City:         model.City,
		Conference:   model.Conference,
		Division:     model.Division,
		FullName:     model.FullName,
		Name:         model.Name,
	}
}

func stringToInt(stringVal string) int {
	intVal, _ := strconv.Atoi(stringVal)
	return intVal
}

func intToString(intVal int) string {
	return strconv.Itoa(intVal)
}
