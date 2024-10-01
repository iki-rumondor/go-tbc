package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Result struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null;unique;size:64"`
	Cluster   byte   `gorm:"not_null;"`
	Type      string `gorm:"not_null;"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	CaseID    uint   `gorm:"not_null;"`
	Case      *Case
}

func (m *Result) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.NewString()
	return nil
}
