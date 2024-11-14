package repository

import (
	"backend-ad-campaign/internal/core"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

func (r *PostgresRepository) CreateCampaign(campaign *core.Campaign) (*core.Campaign, error) {
	query := `INSERT INTO campaigns (name, start_date, end_date, budget, target_views, price_per_view) 
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err := r.DB.QueryRow(query, campaign.Name, campaign.StartDate, campaign.EndDate, campaign.Budget, campaign.TargetViews, campaign.PricePerView).Scan(&campaign.ID)
	if err != nil {
		return nil, fmt.Errorf("Error inserting campaign: %w", err)
	}

	return campaign, nil
}

func (r *PostgresRepository) GetCampaignByID(id string) (*core.Campaign, error) {
	var campaign core.Campaign

	query := `SELECT id, name, start_date, end_date, budget, target_views, price_per_view FROM campaigns WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&campaign.ID, &campaign.Name, &campaign.StartDate, &campaign.EndDate, &campaign.Budget, &campaign.TargetViews, &campaign.PricePerView)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Campaign not found")
		}
		return nil, err
	}

	return &campaign, nil
}

func (r *PostgresRepository) GetAllCampaigns() ([]*core.Campaign, error) {
	query := `SELECT id, name, start_date, end_date, budget, target_views, price_per_view FROM campaigns`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving campaigns: %v", err)
	}
	defer rows.Close()

	var campaigns []*core.Campaign
	for rows.Next() {
		var campaign core.Campaign
		if err := rows.Scan(&campaign.ID, &campaign.Name, &campaign.StartDate, &campaign.EndDate, &campaign.Budget, &campaign.TargetViews, &campaign.PricePerView); err != nil {
			return nil, fmt.Errorf("error scanning campaign: %v", err)
		}
		campaigns = append(campaigns, &campaign)
	}

	return campaigns, nil
}

func (r *PostgresRepository) DeleteCampaign(id string) error {
	query := `DELETE FROM campaigns WHERE id = $1`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error deleting campaign: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Campaign with ID %s not found", id)
	}

	return nil
}

func (r *PostgresRepository) UpdateCampaign(campaign *core.Campaign) (*core.Campaign, error) {
	query := `UPDATE campaigns SET name = $1, start_date = $2, end_date = $3, budget = $4, target_views = $5, price_per_view = $6 WHERE id = $7 RETURNING id, name, start_date, end_date, budget, target_views, price_per_view`
	row := r.DB.QueryRow(query, campaign.Name, campaign.StartDate, campaign.EndDate, campaign.Budget, campaign.TargetViews, campaign.PricePerView, campaign.ID)

	var updatedCampaign core.Campaign
	err := row.Scan(&updatedCampaign.ID, &updatedCampaign.Name, &updatedCampaign.StartDate, &updatedCampaign.EndDate, &updatedCampaign.Budget, &updatedCampaign.TargetViews, &updatedCampaign.PricePerView)
	if err != nil {
		return nil, fmt.Errorf("error updating campaign: %v", err)
	}

	return &updatedCampaign, nil
}
