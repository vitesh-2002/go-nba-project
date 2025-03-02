package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const url = "https://api.balldontlie.io/v1/players?per_page=25&page=1"

// use Get for client functions, Fetch for repo/DB functions
func GetPlayers() ([]byte, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return nil, err
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
	}
	req.Header.Set("Authorization", os.Getenv("BALLDONTLIE_API_KEY"))
	req.Header.Set("Accept", "application/json")

	// Make the GET request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making API request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API returned non-OK status: %s\n%s\n", resp.Status, resp.Body)
		return nil, err
	}

	// Read the response body and return the read in bytes
	return io.ReadAll(resp.Body)
}
