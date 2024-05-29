package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init_env() {
	if err := godotenv.Load(); err != nil {
		fmt.Errorf("error loading .env file: %w", err)
	}
}

func main() {
	init_env() // fikk ikke til å kalle funksjonen bare init, vet ikke hvorfor.
 	mux := http.NewServeMux()

 	mux.Handle("/", &homeHandler{})
	mux.Handle("/search", &searchHandler{})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for /")
 	w.Write([]byte("Hello World!"))
}

type searchHandler struct{}

func (s *searchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		log.Println("Title parameter is missing")
		http.Error(w, "Title parameter is missing", http.StatusBadRequest)
		return
	}

	log.Printf("Received search request for title: %s", title)

	apiKey := os.Getenv("OMDBAPI_KEY")
	if apiKey == "" {
		log.Println("OMDB API key is missing")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s", apiKey, title)
	log.Printf("Fetching data from OMDb API with URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to fetch data from OMDb API: %v", err)
		http.Error(w, "Failed to fetch data from OMDb API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error response from OMDb API: %d", resp.StatusCode)
		http.Error(w, "Error response from OMDb API", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully fetched data from OMDb API: %s", body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Failed to parse JSON response: %v", err)
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

	log.Println("Search request handled successfully")
}