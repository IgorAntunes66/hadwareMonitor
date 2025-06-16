package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RAMStats struct {
	TotalGB     float64 `json:"total_gb"`
	UsedGB      float64 `json:"used_gb"`
	UsedPercent float64 `json:"used_percent"`
}

func main() {
	http.HandleFunc("/", olaHandler)
	http.HandleFunc("stats", statsHandler)

	fmt.Println("Servidor escutando na porta 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}

func olaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor funcionando!")
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	ramStats := RAMStats{TotalGB: 16, UsedGB: 16, UsedPercent: 50}
	jsonBytes, err := json.Marshal(ramStats)
	if err != nil {
		log.Printf("Erro ao converter a struct em JSON: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
