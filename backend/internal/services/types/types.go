package types

import "time"

// Auth
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

// Politicians
type PoliticianResponse struct {
	ID             uint   `json:"id"`
	PoliticianCode string `json:"politician_code"`
	Name           string `json:"name"`
	Party          string `json:"party"`
}

// Campaigns
type CampaignResponse struct {
	ID       uint   `json:"id"`
	Year     int    `json:"year"`
	District string `json:"district"`
}

// Promises
type PromiseResponse struct {
	ID          uint      `json:"id"`
	Detail      string    `json:"detail"`
	AnnouncedAt time.Time `json:"announced_at"`
	Status      string    `json:"status"`

	Politician PoliticianResponse `json:"politician"`
	Campaign   CampaignResponse   `json:"campaign"`
}

type PromiseUpdateResponse struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	Note      string    `json:"note"`
}

type PromiseDetailResponse struct {
	PromiseResponse
	Updates []PromiseUpdateResponse `json:"updates"`
}

type CreatePromiseUpdateRequest struct {
	UpdatedAt time.Time `json:"updated_at"`
	Note      string    `json:"note"`
}
