package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tushargupta7/kong/db"
	"strconv"
)

// GetServices returns a list of services with filtering, sorting, and pagination
func GetServices(c *fiber.Ctx) error {
	search := c.Query("search", "")
	sortBy := c.Query("sort_by", "name")
	order := c.Query("order", "asc")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	offset := (page - 1) * limit
	query := fmt.Sprintf(`
		SELECT id, name, description FROM services
		WHERE name ILIKE '%%%s%%'
		ORDER BY %s %s
		LIMIT %d OFFSET %d`, search, sortBy, order, limit, offset)

	rows, err := db.DB.Query(query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch services"})
	}
	defer rows.Close()

	var services []fiber.Map
	for rows.Next() {
		var id int
		var name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to parse services"})
		}
		services = append(services, fiber.Map{"id": id, "name": name, "description": description})
	}

	return c.JSON(fiber.Map{
		"services": services,
		"page":     page,
		"limit":    limit,
	})
}

// GetService fetches details of a single service by ID
func GetService(c *fiber.Ctx) error {
	id := c.Params("id")

	var name, description string
	err := db.DB.QueryRow("SELECT name, description FROM services WHERE id = $1", id).Scan(&name, &description)
	if err == sql.ErrNoRows {
		return c.Status(404).JSON(fiber.Map{"error": "Service not found"})
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch service"})
	}

	return c.JSON(fiber.Map{"id": id, "name": name, "description": description})
}

// GetServiceVersions fetches all versions of a specific service
func GetServiceVersions(c *fiber.Ctx) error {
	id := c.Params("id")

	rows, err := db.DB.Query("SELECT version_number FROM versions WHERE service_id = $1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch versions"})
	}
	defer rows.Close()

	var versions []string
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to parse versions"})
		}
		versions = append(versions, version)
	}

	return c.JSON(fiber.Map{"service_id": id, "versions": versions})
}
