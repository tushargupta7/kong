package errors

// Error messages constants
const (
	ErrInvalidPayload        = "Invalid request payload"
	ErrInvalidServiceID      = "Invalid service ID"
	ErrFailedToCreateVersion = "Failed to create version"
	ErrFailedToUpdateVersion = "Failed to update version"
	ErrVersionNotFound       = "Version not found"
	ErrFailedToDeleteVersion = "Failed to delete version"
	ErrFailedToFetchVersions = "Failed to fetch versions"
	ErrServiceNotFound       = "Service not found"
	ErrServiceCreation       = "Failed to create service"
	ErrServiceUpdate         = "Failed to update service"
	ErrServiceDeletion       = "Failed to delete service"
	ErrFetchingService       = "Failed to fetch service"
	ErrFetchingServices      = "Failed to fetch services"
	ErrConnectDatabase       = "Failed to connect to DB"
)
