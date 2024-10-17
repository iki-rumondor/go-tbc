package migrate

import (
	"fmt"
	"log"

	"github.com/iki-rumondor/go-tbc/internal/app/structs/models"
	"github.com/iki-rumondor/go-tbc/internal/config"
	"gorm.io/gorm"
)

func ReadTerminal(db *gorm.DB, args []string) {
	switch {
	case args[1] == "fresh":
		if err := freshDatabase(db); err != nil {
			log.Fatal(err.Error())
		}
	case args[1] == "migrate":
		if err := migrateDatabase(db); err != nil {
			log.Fatal(err.Error())
		}
	default:
		fmt.Println("Hello, Nice to meet you")
	}
}

func freshDatabase(db *gorm.DB) error {
	for _, model := range GetAllModels() {
		if err := db.Migrator().DropTable(model.Model); err != nil {
			return err
		}
	}

	for _, model := range GetAllModels() {
		if err := db.Debug().AutoMigrate(model.Model); err != nil {
			return err
		}
	}

	for _, role := range config.SYSTEM_ROLES {
		db.Create(&models.Role{
			Name: role,
		})
	}

	db.Create(&config.ADMIN_USER)
	db.Create(&config.KADIS_USER)

	return nil
}

func migrateDatabase(db *gorm.DB) error {
	for _, model := range GetAllModels() {
		if err := db.Debug().AutoMigrate(model.Model); err != nil {
			return err
		}
	}
	return nil
}
