package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Hardcoded character ID
	charID := uint32(28293967)

	s := godestone.NewScraper(bingode.New(), godestone.EN)
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
