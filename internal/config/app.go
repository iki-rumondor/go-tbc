package config

import (
	"github.com/iki-rumondor/go-tbc/internal/app/layers/handlers"
	"github.com/iki-rumondor/go-tbc/internal/app/layers/repositories"
	"github.com/iki-rumondor/go-tbc/internal/app/layers/services"
	"gorm.io/gorm"
)

type Handlers struct {
	AuthHandler       *handlers.AuthHandler
	ManagementHandler *handlers.ManagementHandler
	FetchHandler      *handlers.FetchHandler
	ProcessingHandler *handlers.ProcessingHandler
}

func GetAppHandlers(db *gorm.DB) *Handlers {

	auth_repo := repositories.NewAuthInterface(db)
	auth_service := services.NewAuthService(auth_repo)
	auth_handler := handlers.NewAuthHandler(auth_service)

	management_repo := repositories.NewManagementInterface(db)
	management_service := services.NewManagementService(management_repo)
	management_handler := handlers.NewManagementHandler(management_service)

	fetch_repo := repositories.NewFetchInterface(db)
	fetch_service := services.NewFetchService(fetch_repo)
	fetch_handler := handlers.NewFetchHandler(fetch_service)

	processing_repo := repositories.NewProcessingInterface(db)
	processing_service := services.NewProcessingService(processing_repo)
	processing_handler := handlers.NewProcessingHandler(processing_service)

	return &Handlers{
		AuthHandler:       auth_handler,
		ManagementHandler: management_handler,
		FetchHandler:      fetch_handler,
		ProcessingHandler: processing_handler,
	}
}
