package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

// Handler is the entry point for Vercel serverless
func Handler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	s := godestone.NewScraper(bingode.New(), godestone.EN)

	c, err := s.FetchCharacter(uint32(id))
	if err != nil {
		http.Error(w, "Error fetching character", http.StatusInternalServerError)
		log.Println("FetchCharacter error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.MarshalIndent(c, "", "  ")
	w.Write(jsonData)
}
