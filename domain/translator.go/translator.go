package translator

import (
	"nba-players/models/player"
	"nba-players/models/team"
	"nba-players/representations"
)

// converting a rep to team and player models
// iterate through each player in []Data
// translate player rep to player model
// team rep will automatically be translated to team model
// this function returns []player.Model for now.
// TODO: return []team.Model as well w/o duplicates.
func Translate(player representations.APIResponse) {
}

// function to convert a player rep into a domain object
func ToPlayerModel(representations.Player) player.Model {
	return player.Model{}
}

// function to convert a team rep into a domain object
func ToTeamModel(representations.Team) team.Model {
	return team.Model{}
}
