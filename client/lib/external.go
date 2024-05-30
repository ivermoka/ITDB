package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

func HandleRegister(args string) (string, error) {
	credentials := strings.Fields(args)

	if len(credentials) != 3 {
		return "", errors.New("invalid number of credentials")
	}
	username := credentials[0]
	mail := credentials[1]
	password := credentials[2]

	url := fmt.Sprintf("http://localhost:8080/auth/register?username=%s&mail=%s&password=%s", username, mail, password)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Failed to register user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Register API responded with status: %s", resp.Status)
	}

	return fmt.Sprintf("Registered user %s. Credentials:\nUsername=%s\nMail=%s\nPassword=%s", username, username, mail, password), nil
}

func HandleLogin(args string) (string, error) {
	credentials := strings.Fields(args)

	if len(credentials) != 2 {
		return "", errors.New("invalid number of credentials")
	}
	username := credentials[0]
	password := credentials[1]

	url := fmt.Sprintf("http://localhost:8080/auth/login?username=%s&password=%s", username, password)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Failed to login user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Login API responded with status: %s", resp.Status)
	}

	return fmt.Sprintf("Logged in user %s. Credentials:\nUsername=%s\nPassword=%s", username, username, password), nil
}