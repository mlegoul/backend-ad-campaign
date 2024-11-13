package core

type Campaign struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Budget      float64 `json:"budget"`
	TargetViews int     `json:"target_views"`
}
