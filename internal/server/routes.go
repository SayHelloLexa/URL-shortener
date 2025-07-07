package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) configureRoutes() {
	s.router.Use(JsonHeaderMiddleware)

	s.router.HandleFunc("/shorten", s.shortenUrlHandler).Methods("POST")
	s.router.HandleFunc("/longer", s.longUrlHandler).Methods("POST")
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

	fullShortUrl := "http://" + s.config.Domain + "/" + shortUrl

	response := struct {
		ShortUrl string `json:"short_url"`
	}{
		ShortUrl: fullShortUrl,
	}

	json.NewEncoder(w).Encode(response)
}

func (s *Server) longUrlHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ShortUrl string `json:"short_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	longUrl, err := s.service.GetLongUrl(request.ShortUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := struct {
		LongUrl string `json:"long_url"`
	}{
		LongUrl: longUrl,
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
