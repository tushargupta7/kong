package dtos

import "time"

// CreateVersionRequest defines the payload for creating a new version
type CreateVersionRequest struct {
	Version      string `json:"version" validate:"required"`
	ReleaseNotes string `json:"release_notes"`
}

// UpdateVersionRequest defines the payload for updating an existing version
type UpdateVersionRequest struct {
	Version      string `json:"version" validate:"required"`
	ReleaseNotes string `json:"release_notes"`
}

// VersionResponse defines the structure of a version response
type VersionResponse struct {
	ID           uint      `json:"id"`
	ServiceID    uint      `json:"service_id"`
	Version      string    `json:"version"`
	ReleaseNotes string    `json:"release_notes"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// PaginationResponse defines a generic pagination response structure
type PaginationResponse struct {
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
	Results interface{} `json:"results"`
}
