package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ScaryFrogg/grim_soul_calculator/internal/common"
)

func (s *Server) TradesForIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "0" {
		http.Error(w, "Id cant be nil/0", http.StatusNotFound)
		return
	}
	writeJSONResponse(w, s.db.GetTradesForItem(id), http.StatusOK)
}

func (s *Server) TradesHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(s.db.GetTrades())
	if err != nil {
		http.Error(w, "Failed to marshal trades data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) DesingnsHandler(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, s.db.GetDesigns(), http.StatusOK)
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

func (s *Server) GetItemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetItemHandel: for Body -> %v, Path if %v", r.Body, r.URL)
	id := r.PathValue("id")
	fmt.Printf("GetItemHandel: for id -> %v", id)
	if id == "0" || id == "" {
		http.Error(w, "Id cant be nil/0", http.StatusNotFound)
		return
	}
	//get design
	materialName, err := s.db.GetItem(id)
	if err != nil {
		fmt.Printf("Failed to get material with Id [%v] design data-> error %v", id, err)
		return
	}

	//get requirements
	designs := s.db.GetDesignsForItem(id)
	w.Header().Set("Content-Type", "application/json")
	info := types.MaterialInfo{Name: materialName, Designs: designs}

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
