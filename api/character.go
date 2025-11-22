package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// --- CORS ---
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	// -------------

	id := uint32(28293967)

	s := godestone.NewScraper(bingode.New(), godestone.EN)
	c, err := s.FetchCharacter(id)
	if err != nil {
		http.Error(w, "Error fetching character", http.StatusInternalServerError)
		log.Println("FetchCharacter error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.MarshalIndent(c, "", "  ")
	w.Write(jsonData)
}
