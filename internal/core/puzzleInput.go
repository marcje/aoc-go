package core

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// GetDailyInput fetches the daily input of the given year and day.
// AOC_SESSION_COOKIE must be set as an environment variable.
func GetDailyInput(year, day int) ([]string, error) {
	inputUrl := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	sessionCookie := os.Getenv("AOC_SESSION_COOKIE")

	if sessionCookie == "" {
		return nil, fmt.Errorf("AOC_SESSION_COOKIE environment variable is not defined")
	}

	req, err := http.NewRequest("GET", inputUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error initializing HTTP request: %w", err)
	}

	client := &http.Client{}
	req.Header.Set("Cookie", "session="+sessionCookie)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading page: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(bodyBytes)), "\n")
	return lines, nil
}
