package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tushargupta7/kong/dtos"
	"github.com/tushargupta7/kong/errors"
	"github.com/tushargupta7/kong/repositories"
)

// Constants for error messages and defaults (unchanged)
const (
	DefaultSortBy    = "name"
	DefaultSortOrder = "asc"
	DefaultPage      = 1
	DefaultLimit     = 10
)

var serviceRepo repositories.ServiceRepository

// Initialize the service repository
func InitServiceHandler(repo repositories.ServiceRepository) {
	serviceRepo = repo
}

// CreateService creates a new service
func CreateService(c *fiber.Ctx) error {
	var req dtos.CreateServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.New(fiber.StatusBadRequest, errors.ErrInvalidPayload, err, nil)
	}

	// Validate DTO
	if req.Name == "" {
		return errors.New(fiber.StatusBadRequest, "Service name is required", nil, nil)
	}

	service, err := serviceRepo.InsertService(req.Name, req.Description)
	if err != nil {
		return errors.New(fiber.StatusInternalServerError, errors.ErrServiceCreation, err, map[string]interface{}{"name": req.Name})
	}

	return c.Status(fiber.StatusCreated).JSON(service)
}

// UpdateService updates an existing service
func UpdateService(c *fiber.Ctx) error {
	var req dtos.UpdateServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.New(fiber.StatusBadRequest, errors.ErrInvalidPayload, err, nil)
	}

	id := c.Params("id")
	idi, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errors.New(fiber.StatusBadRequest, "Invalid service ID", err, nil)
	}
	// Validate if service exists
	existingService, err := serviceRepo.GetServiceByID(idi)
	if err == sql.ErrNoRows {
		return errors.New(fiber.StatusNotFound, errors.ErrServiceNotFound, err, map[string]interface{}{"id": id})
	} else if err != nil {
		return errors.New(fiber.StatusInternalServerError, errors.ErrFetchingService, err, map[string]interface{}{"id": id})
	}

	// Update service
	updatedService, err := serviceRepo.UpdateServiceByID(id, req.Name, req.Description)
	if err != nil {
		return errors.New(fiber.StatusInternalServerError, errors.ErrServiceUpdate, err, map[string]interface{}{"id": id})
	}

	updatedService.ID = existingService.ID // Retain original ID
	return c.JSON(updatedService)
}

// DeleteService deletes a service
func DeleteService(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := serviceRepo.DeleteServiceByID(id); err != nil {
		return errors.New(fiber.StatusInternalServerError, errors.ErrServiceDeletion, err, map[string]interface{}{"id": id})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Service deleted successfully"})
}

// GetServices fetches services with pagination and filtering
func GetServices(c *fiber.Ctx) error {
	search := c.Query("search", "")
	sortBy := c.Query("sort_by", DefaultSortBy)
	order := c.Query("order", DefaultSortOrder)
	page, limit := parsePaginationParams(c)

	offset := (page - 1) * limit

	services, err := serviceRepo.GetServices(search, sortBy, order, limit, offset)
	if err != nil {
		return errors.New(fiber.StatusInternalServerError, errors.ErrFetchingServices, err, nil)
	}

	return c.JSON(dtos.PaginationResponse{
		Page:    page,
		Limit:   limit,
		Results: services,
	})
}

// GetService fetches a single service by ID
func GetService(c *fiber.Ctx) error {
	id := c.Params("id")

	// Convert string ID to int64
	idi, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errors.New(fiber.StatusBadRequest, "Invalid service ID", err, nil)
	}

	service, err := serviceRepo.GetServiceByID(idi)
	if err == sql.ErrNoRows {
		// Create a map variable before passing it to errors.New
		context := map[string]interface{}{"id": id}
		return errors.New(fiber.StatusNotFound, errors.ErrServiceNotFound, err, context)
	} else if err != nil {
		return errors.New(fiber.StatusInternalServerError, errors.ErrFetchingService, err, map[string]interface{}{"id": id})
	}

	return c.JSON(service)
}

// Helper function to parse pagination parameters
func parsePaginationParams(c *fiber.Ctx) (int, int) {
	page, err := strconv.Atoi(c.Query("page", strconv.Itoa(DefaultPage)))
	if nil != err || page < 1 {
		page = DefaultPage
	}

	limit, err := strconv.Atoi(c.Query("limit", strconv.Itoa(DefaultLimit)))
	if nil != err || limit < 1 {
		limit = DefaultLimit
	}

	return page, limit
}
