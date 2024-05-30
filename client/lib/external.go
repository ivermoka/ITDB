package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type User struct {
	Username string
	Jwt 	 string
}

var u User

// func FetchDataFromAPI(title string) (string, error) {
// 	url := fmt.Sprintf("%s/search?title=%s", os.Getenv("VM_IP"), title)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to fetch data from API: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf("API responded with status: %s", resp.Status)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to read response body: %w", err)
// 	}

// 	return string(body), nil
// }

func HandleSearch(args string) (string, error) {
	encodedTitle := url.QueryEscape(args)  // hadde problemer formatet til args, så la jeg til dette, og nå funker det :0
	url := fmt.Sprintf("http://%s/search?title=%s", os.Getenv("VM_IP"), encodedTitle)	
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

func HandleUser() (string) {
	if u.Jwt == "" {
		return fmt.Sprintf("Not logged in.")
	}
	return fmt.Sprintf("User infomation\nUsername:%s\nJWT:%s", u.Username, u.Jwt)
}

func HandleRegister(args string) (string, error) {
	credentials := strings.Fields(args)

	if len(credentials) != 3 {
		return "", errors.New("invalid number of credentials")
	}
	username := credentials[0]
	mail := credentials[1]
	password := credentials[2]

	url := fmt.Sprintf("http://%s/auth/register?username=%s&mail=%s&password=%s", os.Getenv("VM_IP"), username, mail, password)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Failed to register user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Register API responded with status: %s", resp.Status)
	}

	jwt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read response: %w", err)
	}
	u.Jwt = string(jwt)
	u.Username = username

	return fmt.Sprintf("Registered user %s. Credentials:\nUsername=%s\nMail=%s\nPassword=%s\n%s", username, username, mail, password, jwt), nil
}

func HandleLogin(args string) (string, error) {
	credentials := strings.Fields(args)

	if len(credentials) != 2 {
		return "", errors.New("invalid number of credentials")
	}
	username := credentials[0]
	password := credentials[1]

	url := fmt.Sprintf("http://%s//auth/login?username=%s&password=%s", os.Getenv("VM_IP"), username, password)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Failed to login user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Login API responded with status: %s", resp.Status)
	}

	jwt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read response: %w", err)
	}
	u.Jwt = string(jwt)
	u.Username = username

	return fmt.Sprintf("Logged in user %s. Credentials:\nUsername=%s\nPassword=%s\n%s", username, username, password, jwt), nil
}