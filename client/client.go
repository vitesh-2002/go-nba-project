package client

import (
	"encoding/json"
	"fmt"
	"io"
	"nba-players/representations"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const url = "https://api.balldontlie.io/v1/players?per_page=25&page=1"

// use Get for client functions, Fetch for repo/DB functions
func GetPlayers() (representations.APIResponse, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return representations.APIResponse{}, err
	}

	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a request to send to the API
	// Set headers for the request - Authorization, Accept, etc.
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Error creating API request: %v\n", err)
		return representations.APIResponse{}, nil
	}
	req.Header.Set("Authorization", os.Getenv("BALLDONTLIE_API_KEY"))
	req.Header.Set("Accept", "application/json")

	// Make the GET request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making API request: %v\n", err)
		return representations.APIResponse{}, err
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API returned non-OK status: %s\n%s\n", resp.Status, resp.Body)
		return representations.APIResponse{}, err
	}

	// Read the response body and return the read in bytes
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the response body bytes: \n%v\n", err)
		return representations.APIResponse{}, err
	}

	// Unmarshal the JSON into our struct
	var apiResponse representations.APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("Error umarshaling the JSON: \n%v\n", err)
		return apiResponse, err
	}

	return apiResponse, nil
}
