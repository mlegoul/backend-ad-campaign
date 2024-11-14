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

	// Si price_per_view est 0 ou non défini, on le calcule à partir du budget et des affichages
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
	// Récupérer l'ID depuis les paramètres de la route
	vars := mux.Vars(r)
	id := vars["id"]

	// Log de l'ID reçu
	log.Printf("Received request to get campaign with ID: %s", id)

	// Utiliser le repository pour récupérer la campagne
	campaign, err := h.Service.GetCampaignByID(id)
	if err != nil {
		// Log de l'erreur
		log.Printf("Error retrieving campaign: %v", err)
		http.Error(w, "Campaign not found", http.StatusNotFound)
		return
	}

	// Retourner la campagne trouvée en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(campaign)
}
