package repositories

import "github.com/tushargupta7/kong/dtos"

// ServiceRepository defines the methods for service-related database operations
type ServiceRepository interface {
	InsertService(name, description string) (dtos.ServiceResponse, error)
	GetServiceByID(id int64) (dtos.ServiceResponse, error)
	UpdateServiceByID(id string, name, description string) (dtos.ServiceResponse, error)
	DeleteServiceByID(id string) error
	GetServices(search, sortBy, order string, limit, offset int) ([]dtos.ServiceResponse, error)
}
