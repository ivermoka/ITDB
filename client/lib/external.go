package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


func FetchDataFromAPI(title string) (string, error) {
	url := fmt.Sprintf("http://localhost:8080/search?title=%s", title)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch data from API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API responded with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

func HandleSearch(args string) (string, error) {
	encodedTitle := url.QueryEscape(args)  // hadde problemer formatet til args, så la jeg til dette, og nå funker det :0
	url := fmt.Sprintf("http://localhost:8080/search?title=%s", encodedTitle)	
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to perform search: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("search API responded with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read search response: %w", err)
	}
	var omdbResponse OMDbResponse
	if err := json.Unmarshal(body, &omdbResponse); err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return FormatResponse(omdbResponse), nil
}

func HandleReview(args string) (string, error) {
	// TODO: lag review funksjonalitet
	return fmt.Sprintf("Review for %s", args), nil
}