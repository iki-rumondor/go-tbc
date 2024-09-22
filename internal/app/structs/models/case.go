package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Case struct {
	ID             uint   `gorm:"primaryKey"`
	Uuid           string `gorm:"not_null;unique;size:64"`
	Year           string `gorm:"not_null;"`
	ChildCount     int64  `gorm:"not_null;"`
	AdultCount     int64  `gorm:"not_null;"`
	MaleCount      int64  `gorm:"not_null;"`
	FemaleCount    int64  `gorm:"not_null;"`
	CreatedAt      int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt      int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	HealthCenterID uint   `gorm:"not_null;"`
	HealthCenter   *HealthCenter
}

func (m *Case) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
