package interfaces

import "github.com/iki-rumondor/go-tbc/internal/app/structs/models"

type AuthInterface interface {
	FirstUserByUsername(username string) (*models.User, error)
}
