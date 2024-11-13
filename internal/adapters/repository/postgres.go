package repository

import (
	"backend-ad-campaign/internal/core"
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

func (r *PostgresRepository) CreateCampaign(campaign *core.Campaign) (*core.Campaign, error) {
	query := `INSERT INTO campaigns (name, start_date, end_date, budget, target_views) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRow(query, campaign.Name, campaign.StartDate, campaign.EndDate, campaign.Budget, campaign.TargetViews).Scan(&campaign.ID)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}
