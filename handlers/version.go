package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	dto "github.com/tushargupta7/kong/dtos"
	"github.com/tushargupta7/kong/errors"
	"github.com/tushargupta7/kong/repositories"
)

// CreateVersion handler with custom error messages and exceptions
func CreateVersion(c *fiber.Ctx) error {
	var req dto.CreateVersionRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.New(fiber.StatusBadRequest, errors.ErrInvalidPayload, err, nil)
	}

	serviceID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errors.New(fiber.StatusBadRequest, errors.ErrInvalidServiceID, err, map[string]interface{}{"service_id": c.Params("id")})
	}

	version, err := repositories.CreateVersion(serviceID, req)
	if err != nil {
		return errors.New(fiber.StatusInternalServerError, errors.ErrFailedToCreateVersion, err, map[string]interface{}{"service_id": serviceID})
	}

	return c.Status(fiber.StatusCreated).JSON(version)
}

// UpdateVersion handler with exception wrapping
func UpdateVersion(c *fiber.Ctx) error {
	serviceID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errors.New(fiber.StatusBadRequest, errors.ErrInvalidServiceID, err, map[string]interface{}{"service_id": c.Params("id")})
	}

	versionID, err := strconv.Atoi(c.Params("versionId"))
	if err != nil {
		return errors.New(fiber.StatusBadRequest, "Invalid version ID", err, map[string]interface{}{"version_id": c.Params("versionId")})
	}

	var req dto.UpdateVersionRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.New(fiber.StatusBadRequest, errors.ErrInvalidPayload, err, nil)
	}

	err = repositories.UpdateVersion(versionID, serviceID, req)
	if err != nil {
		if err.Error() == "version not found" {
			return errors.New(fiber.StatusNotFound, errors.ErrVersionNotFound, err, map[string]interface{}{"service_id": serviceID, "version_id": versionID})
		}
		return errors.New(fiber.StatusInternalServerError, errors.ErrFailedToUpdateVersion, err, map[string]interface{}{"service_id": serviceID, "version_id": versionID})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Version updated successfully"})
}

// DeleteVersion handler with custom error handling
func DeleteVersion(c *fiber.Ctx) error {
	versionID, err := strconv.Atoi(c.Params("versionId"))
	if err != nil {
		return errors.New(fiber.StatusBadRequest, "Invalid version ID", err, map[string]interface{}{"version_id": c.Params("version_id")})
	}

	err = repositories.DeleteVersion(versionID)
	if err != nil {
		return errors.New(fiber.StatusInternalServerError, errors.ErrFailedToDeleteVersion, err, map[string]interface{}{"version_id": versionID})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Version deleted successfully"})
}

// GetServiceVersions handler with enhanced exception handling
func GetServiceVersions(c *fiber.Ctx) error {
	serviceID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errors.New(fiber.StatusBadRequest, errors.ErrInvalidServiceID, err, map[string]interface{}{"service_id": c.Params("id")})
	}

	versions, err := repositories.GetServiceVersions(serviceID)
	if err != nil {
		return errors.New(fiber.StatusInternalServerError, errors.ErrFailedToFetchVersions, err, map[string]interface{}{"service_id": serviceID})
	}

	return c.JSON(fiber.Map{"service_id": serviceID, "versions": versions})
}
