package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

// Handler is required by Vercel for serverless Go
func Handler(w http.ResponseWriter, r *http.Request) {
	s := godestone.NewScraper(bingode.New(), godestone.EN)

	// Hard-coded character ID
	charID := uint32(28293967)

	c, err := s.FetchCharacter(charID)
	if err != nil {
		http.Error(w, "Error fetching character", http.StatusInternalServerError)
		log.Println("FetchCharacter error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.MarshalIndent(c, "", "  ")
	w.Write(jsonData)
}
