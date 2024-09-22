package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HealthCenter struct {
	ID        uint    `gorm:"primaryKey"`
	Uuid      string  `gorm:"not_null;unique;size:64"`
	Name      string  `gorm:"not_null;size:128"`
	Longitude float64 `gorm:"not_null;"`
	Latitude  float64 `gorm:"not_null;"`
	CreatedAt int64   `gorm:"autoCreateTime:milli"`
	UpdatedAt int64   `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}

func (m *HealthCenter) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}