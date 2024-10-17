package config

import "github.com/iki-rumondor/go-tbc/internal/app/structs/models"

var SYSTEM_ROLES = []string{"ADMIN", "KADIS"}
var ADMIN_USER = models.User{
	Name:     "Administrator",
	Username: "admin",
	Password: "admin123",
	Active:   true,
	RoleID:   1,
}

var KADIS_USER = models.User{
	Name:     "Kepala Dinas",
	Username: "kadis",
	Password: "kadis123",
	Active:   true,
	RoleID:   2,
}
