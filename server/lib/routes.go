package lib

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


type homeHandler struct{}
type searchHandler struct{}
type registerHandler struct {
	db *sql.DB
}
type loginHandler struct {
	db *sql.DB
}
type reviewHandler struct{
	db *sql.DB
}

func Server() {
	db := Init()
 	mux := http.NewServeMux()

 	mux.Handle("/", &homeHandler{})
	mux.Handle("/search", &searchHandler{})
	mux.Handle("/auth/register", &registerHandler{db})
	mux.Handle("/auth/login", &loginHandler{db})
	mux.Handle("/protected/review", &reviewHandler{db})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}



func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for /")
 	w.Write([]byte("Hello World!"))
}

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

func (s *registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	mail := r.URL.Query().Get("mail")
	password := r.URL.Query().Get("password")

	if username == "" || mail == "" || password == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	jwt, err := createUser(s.db, username, mail, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jwt))
}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if username == "" || password == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	jwt, err := loginUser(h.db, username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jwt))
}

func (h *reviewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]
	
	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}
	
	fmt.Fprint(w, "Welcome to the review page")
}