package api

import (
	"backend-ad-campaign/internal/core"
	"backend-ad-campaign/internal/ports"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type CampaignHandler struct {
	Service ports.Repository
}

func NewCampaignHandler(service ports.Repository) *CampaignHandler {
	return &CampaignHandler{Service: service}
}

func (h *CampaignHandler) HandleCreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign core.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err) // Log l'erreur ici
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("Received campaign: %+v", campaign)

	if campaign.PricePerView == 0 {
		campaign.PricePerView = campaign.Budget / float64(campaign.TargetViews)
	}

	createdCampaign, err := h.Service.CreateCampaign(&campaign)
	if err != nil {
		http.Error(w, "Error creating campaign", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdCampaign)
}

func (h *CampaignHandler) HandleGetCampaignByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Printf("Received request to get campaign with ID: %s", id)

	campaign, err := h.Service.GetCampaignByID(id)
	if err != nil {
		log.Printf("Error retrieving campaign: %v", err)
		http.Error(w, "Campaign not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(campaign)
}

func (h *CampaignHandler) HandleGetAllCampaigns(w http.ResponseWriter, r *http.Request) {
	campaigns, err := h.Service.GetAllCampaigns()
	if err != nil {
		http.Error(w, "Error retrieving campaigns", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(campaigns)
}

func (h *CampaignHandler) HandleDeleteCampaign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.Service.DeleteCampaign(id)
	if err != nil {
		http.Error(w, "Campaign not found or unable to delete", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CampaignHandler) HandleUpdateCampaign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var campaign core.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	campaign.ID = id

	updatedCampaign, err := h.Service.UpdateCampaign(&campaign)
	if err != nil {
		http.Error(w, "Error updating campaign", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCampaign)
}
