package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	domain = "localhost:8080"
)

func (s *Server) configureRoutes() {
	s.router.Use(JsonHeaderMiddleware)

	s.router.HandleFunc("/shorten", s.shortenUrlHandler).Methods("POST")
	s.router.HandleFunc("/{shortUrl}", s.redirectHandler).Methods("GET")
}

func (s *Server) shortenUrlHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		LongUrl string `json:"long_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	shortUrl, err := s.service.ShortenUrl(request.LongUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fullShortUrl := "http://" + domain + "/" + shortUrl

	response := struct {
		ShortUrl string `json:"short_url"`
	}{
		ShortUrl: fullShortUrl,
	}

	json.NewEncoder(w).Encode(response)
}

func (s *Server) redirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]

	longUrl, err := s.service.GetLongUrl(shortUrl)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
	}

	http.Redirect(w, r, longUrl, http.StatusFound)
}
