package main

import (
	"encoding/json"
	"fmt"
	"nba-players/client"
	"nba-players/errorhandler"
	"nba-players/representations"
)

var errors []error

func main() {
	// Send request to get player information
	body, err := client.GetPlayers()
	if err != nil {
		errors = append(errors, err)
		errorhandler.Return(errors)
		return
	}

	// TODO: move unmarshaling json, translating rep and domain structs into service layer

	// Unmarshal the JSON into our struct
	var apiResponse representations.APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		errors = append(errors, err)
		errorhandler.Return(errors)
		return
	}

	// TODO: eventually just print out all the data into a file or terminal
	// TODO: eventually hook this up to a DB connection and store players in a database table
	// 			- need to set up gorm tags, domain structs
	// 			- ensure db connection works, fetch records from DB when displaying players/teams information

	// Print some sample data
	fmt.Printf("Total players found: %d\n", apiResponse.Meta.TotalCount)
	fmt.Printf("Current page: %d of %d\n", apiResponse.Meta.CurrentPage, apiResponse.Meta.TotalPages)
	fmt.Println("\nSample of first 5 players:")

	// Limit to first 5 players or total if less than 5
	sampleSize := min(5, len(apiResponse.Data))
	for i := 0; i < sampleSize; i++ {
		player := apiResponse.Data[i]
		height := "N/A"
		if player.HeightFeet != nil && player.HeightInches != nil {
			height = fmt.Sprintf("%d'%d\"", *player.HeightFeet, *player.HeightInches)
		}
		weight := "N/A"
		if player.WeightPounds != nil {
			weight = fmt.Sprintf("%d lbs", *player.WeightPounds)
		}

		fmt.Printf("\nPlayer #%d:\n", i+1)
		fmt.Printf("Name: %s %s\n", player.FirstName, player.LastName)
		fmt.Printf("Position: %s\n", player.Position)
		fmt.Printf("Height: %s\n", height)
		fmt.Printf("Weight: %s\n", weight)
		fmt.Printf("Team: %s (%s)\n", player.Team.FullName, player.Team.Abbreviation)
	}
}

// Helper function to find minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
