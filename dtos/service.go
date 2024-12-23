package dtos

import (
	"time"
)

// CreateServiceRequest defines the payload for creating a new service
type CreateServiceRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"max=255"`
}

// UpdateServiceRequest defines the payload for updating an existing service
type UpdateServiceRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"max=255"`
}

// ServiceResponse defines the structure of a service response
type ServiceResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	VersionCount int       `json:"version_count"`
}
