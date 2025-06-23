package handlers

import "github.com/benidevo/website/internal/services"

// Handlers bundles all HTTP handlers
type Handlers struct {
	HomeHandler *HomeHandler
}

// SetupHandlers initializes and returns all HTTP handlers with their service dependencies
func SetupHandlers(services *services.Services) *Handlers {
	return &Handlers{
		HomeHandler: NewHomeHandler(services.ProjectService),
	}
}
