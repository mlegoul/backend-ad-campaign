package core

import (
	"encoding/json"
	"time"
)

type Campaign struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	StartDate    time.Time `json:"start_date,omitempty"`
	EndDate      time.Time `json:"end_date,omitempty"`
	Budget       float64   `json:"budget"`
	TargetViews  int       `json:"target_views"`
	PricePerView float64   `json:"price_per_view"`
}

func (c *Campaign) UnmarshalJSON(data []byte) error {
	// Unmarshal the JSON with custom date format
	type Alias Campaign
	aux := &struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse the custom date format
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, aux.StartDate)
	if err != nil {
		return err
	}
	c.StartDate = startDate

	endDate, err := time.Parse(layout, aux.EndDate)
	if err != nil {
		return err
	}
	c.EndDate = endDate

	return nil
}
