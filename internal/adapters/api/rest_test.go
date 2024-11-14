package api

import (
	"backend-ad-campaign/internal/core"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockRepository struct {
	campaigns map[string]core.Campaign
}

func (m *MockRepository) GetCampaignByID(id string) (*core.Campaign, error) {
	if campaign, exists := m.campaigns[id]; exists {
		return &campaign, nil
	}
	return nil, errors.New("campaign not found")
}

func (m *MockRepository) CreateCampaign(campaign *core.Campaign) (*core.Campaign, error) {
	return campaign, nil
}

func (m *MockRepository) GetAllCampaigns() ([]*core.Campaign, error) {
	var campaigns []*core.Campaign
	for _, campaign := range m.campaigns {
		c := campaign
		campaigns = append(campaigns, &c)
	}
	return campaigns, nil
}

func (m *MockRepository) DeleteCampaign(id string) error {
	if _, exists := m.campaigns[id]; exists {
		delete(m.campaigns, id)
		return nil
	}
	return errors.New("campaign not found")
}

func TestHandleGetCampaignByID_Success(t *testing.T) {
	mockRepo := &MockRepository{
		campaigns: map[string]core.Campaign{
			"1": {ID: "1", Name: "Test Campaign"},
		},
	}
	handler := &CampaignHandler{Service: mockRepo}
	req, err := http.NewRequest("GET", "/campaigns/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	handler.HandleGetCampaignByID(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var campaign core.Campaign
	json.NewDecoder(rr.Body).Decode(&campaign)
	assert.Equal(t, "Test Campaign", campaign.Name)
}

func TestHandleGetCampaignByID_NotFound(t *testing.T) {
	mockRepo := &MockRepository{
		campaigns: map[string]core.Campaign{},
	}
	handler := &CampaignHandler{Service: mockRepo}
	req, err := http.NewRequest("GET", "/campaigns/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "2"})
	rr := httptest.NewRecorder()
	handler.HandleGetCampaignByID(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), "Campaign not found")
}

func TestHandleDeleteCampaign(t *testing.T) {
	mockRepo := &MockRepository{
		campaigns: map[string]core.Campaign{
			"1": {ID: "1", Name: "Campaign to Delete"},
		},
	}
	handler := &CampaignHandler{Service: mockRepo}
	req, err := http.NewRequest("DELETE", "/campaigns/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	handler.HandleDeleteCampaign(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.NotContains(t, mockRepo.campaigns, "1")

	req, err = http.NewRequest("DELETE", "/campaigns/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "2"})
	rr = httptest.NewRecorder()
	handler.HandleDeleteCampaign(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), "Campaign not found or unable to delete")
}

func TestHandleGetAllCampaigns(t *testing.T) {
	type CampaignNoDates struct {
		ID           string
		Name         string
		Budget       int
		TargetViews  int
		PricePerView int
	}
	mockRepo := &MockRepository{
		campaigns: map[string]core.Campaign{
			"1": {ID: "1", Name: "Campaign One", Budget: 1000, TargetViews: 5000, PricePerView: 200},
			"2": {ID: "2", Name: "Campaign Two", Budget: 2000, TargetViews: 10000, PricePerView: 300},
			"3": {ID: "3", Name: "Campaign Three", Budget: 1500, TargetViews: 7500, PricePerView: 250},
		},
	}
	handler := &CampaignHandler{Service: mockRepo}
	req, err := http.NewRequest("GET", "/campaigns", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.HandleGetAllCampaigns(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var campaigns []CampaignNoDates
	err = json.NewDecoder(rr.Body).Decode(&campaigns)
	assert.NoError(t, err)
	assert.Len(t, campaigns, 3)
	assert.Equal(t, "Campaign One", campaigns[0].Name)
	assert.Equal(t, "Campaign Two", campaigns[1].Name)
	assert.Equal(t, "Campaign Three", campaigns[2].Name)
}
