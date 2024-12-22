package repositories

import (
	"errors"
	"github.com/tushargupta7/kong/database"
	dto "github.com/tushargupta7/kong/dtos"
)

// CreateVersion inserts a new version into the database
func CreateVersion(serviceID int, req dto.CreateVersionRequest) (dto.VersionResponse, error) {
	query := `
		INSERT INTO versions (service_id, version_number, release_notes, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id, created_at, updated_at`

	var version dto.VersionResponse
	err := database.DB.QueryRow(query, serviceID, req.Version, req.ReleaseNotes).
		Scan(&version.ID, &version.CreatedAt, &version.UpdatedAt)

	if err != nil {
		return version, err
	}

	version.ServiceID = uint(serviceID)
	version.Version = req.Version
	version.ReleaseNotes = req.ReleaseNotes

	return version, nil
}

// UpdateVersion updates an existing version in the database
func UpdateVersion(versionID int, serviceID int, req dto.UpdateVersionRequest) error {
	checkQuery := `SELECT COUNT(*) FROM versions WHERE id = $1 AND service_id = $2`
	var count int
	err := database.DB.QueryRow(checkQuery, versionID, serviceID).Scan(&count)
	if err != nil || count == 0 {
		return errors.New("version not found")
	}

	updateQuery := `
		UPDATE versions
		SET version_number = $1, release_notes = $2, updated_at = NOW()
		WHERE id = $3 AND service_id = $4`

	_, err = database.DB.Exec(updateQuery, req.Version, req.ReleaseNotes, versionID, serviceID)
	return err
}

// DeleteVersion deletes a version from the database
func DeleteVersion(versionID int) error {
	query := `DELETE FROM versions WHERE id = $1`
	_, err := database.DB.Exec(query, versionID)
	return err
}

// GetServiceVersions fetches all versions for a specific service
func GetServiceVersions(serviceID int) ([]string, error) {
	query := `SELECT version_number FROM versions WHERE service_id = $1`
	rows, err := database.DB.Query(query, serviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var versions []string
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		versions = append(versions, version)
	}

	return versions, nil
}
