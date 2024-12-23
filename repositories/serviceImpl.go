package repositories

import (
	"database/sql"
	"fmt"
	"github.com/tushargupta7/kong/dtos"
	"strconv"
)

type serviceRepositoryImpl struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) ServiceRepository {
	return &serviceRepositoryImpl{db: db}
}

func (r *serviceRepositoryImpl) InsertService(name, description string) (dtos.ServiceResponse, error) {
	query := "INSERT INTO services (name, description, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id, created_at, updated_at"
	var service dtos.ServiceResponse
	err := r.db.QueryRow(query, name, description).Scan(&service.ID, &service.CreatedAt, &service.UpdatedAt)
	if err != nil {
		return dtos.ServiceResponse{}, err
	}
	service.Name = name
	service.Description = description
	return service, nil
}

func (r *serviceRepositoryImpl) GetServiceByID(id int64) (dtos.ServiceResponse, error) {
	var service dtos.ServiceResponse
	query := `SELECT s.id, s.name, s.description, s.created_at, s.updated_at,
	          (SELECT COUNT(*) FROM versions v WHERE v.service_id = s.id) AS version_count
	          FROM services s
	          WHERE s.id = $1`
	err := r.db.QueryRow(query, id).Scan(&service.ID, &service.Name, &service.Description, &service.CreatedAt, &service.UpdatedAt, &service.VersionCount)
	if err != nil {
		return dtos.ServiceResponse{}, err
	}
	return service, nil
}

func (r *serviceRepositoryImpl) UpdateServiceByID(id string, name, description string) (dtos.ServiceResponse, error) {
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return dtos.ServiceResponse{}, fmt.Errorf("invalid ID format: %w", err)
	}

	query := "UPDATE services SET name = $1, description = $2, updated_at = NOW() WHERE id = $3 RETURNING created_at, updated_at"
	var updated dtos.ServiceResponse
	err = r.db.QueryRow(query, name, description, parsedID).Scan(&updated.CreatedAt, &updated.UpdatedAt)
	if err != nil {
		return dtos.ServiceResponse{}, err
	}

	updated.ID = uint(parsedID)
	updated.Name = name
	updated.Description = description
	return updated, nil
}

func (r *serviceRepositoryImpl) DeleteServiceByID(id string) error {
	query := `DELETE FROM services WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *serviceRepositoryImpl) GetServices(search, sortBy, order string, limit, offset int) ([]dtos.ServiceResponse, error) {
	allowedSortBy := map[string]bool{"name": true, "created_at": true, "updated_at": true}
	allowedOrder := map[string]bool{"asc": true, "desc": true}

	if !allowedSortBy[sortBy] {
		sortBy = "name"
	}
	if !allowedOrder[order] {
		order = "asc"
	}

	query := `SELECT s.id, s.name, s.description, s.created_at, s.updated_at, 
          (SELECT COUNT(*) FROM versions v WHERE v.service_id = s.id) AS version_count
          FROM services s
          WHERE s.name ILIKE $1
          ORDER BY ` + sortBy + ` ` + order + `
          LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(query, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []dtos.ServiceResponse
	for rows.Next() {
		var service dtos.ServiceResponse
		if err := rows.Scan(&service.ID, &service.Name, &service.Description, &service.CreatedAt, &service.UpdatedAt, &service.VersionCount); err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}
