package handler

import (
	"io"
	"log"
	"net/http"
	"strconv"
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
	// --------------
	id := 28293967

	resp, err := http.Get(
		"http://localhost:8080/Character/" + strconv.Itoa(id) + "?data=CJ",
	)
	if err != nil {
		http.Error(w, "Error contacting Nodestone", http.StatusInternalServerError)
		log.Println("Nodestone error:", err)
		return
	}
	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading Nodestone response", http.StatusInternalServerError)
		log.Println("Read error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
