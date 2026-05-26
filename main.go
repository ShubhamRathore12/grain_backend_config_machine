package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"

	"machine-config-service/config"
)

func main() {
	// Feature flag: SERVICE_ENABLED controls whether this service starts
	enabled := os.Getenv("SERVICE_ENABLED")
	if strings.ToLower(enabled) == "false" || enabled == "0" {
		log.Println("⚠️  Machine Config Service is DISABLED via SERVICE_ENABLED flag. Exiting.")
		os.Exit(0)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()

	// CORS middleware
	r.Use(corsMiddleware)

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Get all machines (list with metadata)
	api.HandleFunc("/machines", getAllMachines).Methods("GET", "OPTIONS")

	// Get single machine full config
	api.HandleFunc("/machines/{name}", getMachineByName).Methods("GET", "OPTIONS")

	// Get machine config by section
	api.HandleFunc("/machines/{name}/auto", getMachineAutoConfig).Methods("GET", "OPTIONS")
	api.HandleFunc("/machines/{name}/auto-grain", getMachineGrainConfig).Methods("GET", "OPTIONS")
	api.HandleFunc("/machines/{name}/auto-paddy", getMachinePaddyConfig).Methods("GET", "OPTIONS")
	api.HandleFunc("/machines/{name}/outputs", getMachineOutputs).Methods("GET", "OPTIONS")
	api.HandleFunc("/machines/{name}/analog", getMachineAnalog).Methods("GET", "OPTIONS")
	api.HandleFunc("/machines/{name}/tags", getMachineTags).Methods("GET", "OPTIONS")
	api.HandleFunc("/machines/{name}/menu", getMachineMenu).Methods("GET", "OPTIONS")

	// Health check
	api.HandleFunc("/health", healthCheck).Methods("GET")

	log.Printf("🚀 Machine Config Service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"service": "machine-config-service",
	})
}

func getAllMachines(w http.ResponseWriter, r *http.Request) {
	machines := config.GetAllMachines()
	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success":  true,
		"count":    len(machines),
		"machines": machines,
	})
}

func getMachineByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	machine, found := config.GetMachineByName(name)
	if !found {
		respondError(w, http.StatusNotFound, fmt.Sprintf("Machine '%s' not found", name))
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"machine": machine,
	})
}

func getMachineAutoConfig(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	cfg, found := config.GetAutoConfig(name)
	if !found {
		respondError(w, http.StatusNotFound, fmt.Sprintf("Auto config for '%s' not found", name))
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"config":  cfg,
	})
}

func getMachineGrainConfig(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	cfg, found := config.GetGrainConfig(name)
	if !found {
		respondError(w, http.StatusNotFound, fmt.Sprintf("Grain config for '%s' not found", name))
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"config":  cfg,
	})
}

func getMachinePaddyConfig(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	cfg, found := config.GetPaddyConfig(name)
	if !found {
		respondError(w, http.StatusNotFound, fmt.Sprintf("Paddy config for '%s' not found", name))
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"config":  cfg,
	})
}

func getMachineOutputs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	outputs, found := config.GetOutputsConfig(name)
	if !found {
		respondError(w, http.StatusNotFound, fmt.Sprintf("Outputs config for '%s' not found", name))
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"outputs": outputs,
	})
}

func getMachineAnalog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	analog, found := config.GetAnalogConfig(name)
	if !found {
		respondError(w, http.StatusNotFound, fmt.Sprintf("Analog config for '%s' not found", name))
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"analog":  analog,
	})
}

func getMachineTags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	tags, found := config.GetMachineTags(name)
	if !found {
		respondError(w, http.StatusNotFound, fmt.Sprintf("Tags for '%s' not found", name))
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"tags":    tags,
	})
}

func getMachineMenu(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	_ = strings.ToLower(name) // normalize if needed

	menu, found := config.GetMenuConfig(name)
	if !found {
		respondError(w, http.StatusNotFound, fmt.Sprintf("Menu config for '%s' not found", name))
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"menu":    menu,
	})
}
