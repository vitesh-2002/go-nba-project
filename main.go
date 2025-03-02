package main

import (
	"encoding/json"
	"fmt"
	"io"
	"nba-players/representations"
	"net/http"
	"time"
)

func main() {
	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// API endpoint for all players (first page, 100 players per page)
	url := "https://www.balldontlie.io/api/v1/players?per_page=100"

	// Set headers for the request - Authorization, Accept, etc.
	// Get API Key from balldontlie.com

	// Make the GET request
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error making API request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API returned non-OK status: %s\n", resp.Status)
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Unmarshal the JSON into our struct
	var apiResponse representations.APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
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
