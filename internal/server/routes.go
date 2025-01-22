package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ScaryFrogg/grim_soul_calculator/internal/common"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("GET /health", s.healthHandler)
	mux.HandleFunc("GET /weapons", s.WeaponDataHandler)
	mux.HandleFunc("GET /designs", s.DesingnsHandler)
	mux.HandleFunc("GET /design/{id}", s.GetBlueprintHandler)
	mux.HandleFunc("GET /material/{id}", s.GetMaterialHandler)
	mux.HandleFunc("GET /enemies", s.GetEnemiesHandeler)
	mux.HandleFunc("GET /requirement/{id}", s.RequirementHandler)

	// Wrap the mux with CORS middleware
	return s.corsMiddleware(mux)
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

func (s *Server) DesingnsHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(s.db.GetDesigns())
	if err != nil {
		http.Error(w, "Failed to marshal weapon data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) RequirementHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	resp, err := json.Marshal(s.db.GetRequirements(id))
	if err != nil {
		http.Error(w, "Failed to marshal weapon data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) GetMaterialHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "0" {
		http.Error(w, "Id cant be nil/0", http.StatusNotFound)
		return
	}
	//get design
	materialName, err := s.db.GetMaterial(id)
	if err != nil {
		fmt.Printf("Failed to get material with Id [%v] design data-> error %v", id, err)
		return
	}

	//get requirements
	designs := s.db.GetDesignsForMaterial(id)
	w.Header().Set("Content-Type", "application/json")
	info := dto.MaterialInfo{Name: materialName, Designs: designs}

	resp, err := json.Marshal(info)
	if err != nil {
		http.Error(w, "Failed to marshal design data", http.StatusInternalServerError)
		return
	}
	//write response
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}

}

func (s *Server) GetBlueprintHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "0" {
		http.Error(w, "Id cant be nil/0", http.StatusNotFound)
		return
	}
	//get design
	design := s.db.GetDesign(id)

	//get requirements
	requirements := s.db.GetRequirements(id)
	design.Requirements = requirements

	resp, err := json.Marshal(design)
	if err != nil {
		http.Error(w, "Failed to marshal design data", http.StatusInternalServerError)
		return
	}
	//write response
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}

}

func (s *Server) WeaponDataHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(s.db.WeaponData())
	if err != nil {
		http.Error(w, "Failed to marshal weapon data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) GetEnemiesHandeler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(s.db.GetEnemies())
	if err != nil {
		http.Error(w, "Failed to marshal enemies data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(s.db.Health())
	if err != nil {
		http.Error(w, "Failed to marshal health check response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
