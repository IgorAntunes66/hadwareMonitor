package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type RAMStats struct {
	TotalGB     float64 `json:"total_gb"`
	UsedGB      float64 `json:"used_gb"`
	UsedPercent float64 `json:"used_percent"`
}

type CPUStats struct {
	GeneralUsage float64 `json:"general_usage"`
}

type HardwareStats struct {
	CPU CPUStats `json:"cpu"`
	RAM RAMStats `json:"ram"`
}

func main() {
	http.HandleFunc("/stats", statsHandler)

	fmt.Println("Servidor escutando na porta 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	cpu, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Printf("Erro ao obter os dados da CPU: %v", err)
		http.Error(w, "Erro ao obter dados da CPU", http.StatusInternalServerError)
		return
	}

	cpuStats := CPUStats{
		GeneralUsage: cpu[0],
	}

	vMem, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Erro ao obter os dados da RAM: %v", err)
		http.Error(w, "Erro ao obter dados da RAM", http.StatusInternalServerError)
		return
	}

	ramStats := RAMStats{
		TotalGB:     float64(vMem.Total) / 1024 / 1024 / 1024,
		UsedGB:      float64(vMem.Used) / 1024 / 1024 / 1024,
		UsedPercent: vMem.UsedPercent,
	}

	hardwareStats := HardwareStats{
		CPU: cpuStats,
		RAM: ramStats,
	}

	jsonBytes, err := json.Marshal(hardwareStats)
	if err != nil {
		log.Printf("Erro ao converter a struct em JSON: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
