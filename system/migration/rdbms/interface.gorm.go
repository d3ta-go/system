package rdbms

import (
	"github.com/d3ta-go/system/system/handler"
	"gorm.io/gorm"
)

// IGormMigrateRunner represent GormMigrateRunner Interface
type IGormMigrateRunner interface {
	GetID() string
	Run(h *handler.Handler, dbGorm *gorm.DB) error
	RollBack(h *handler.Handler, dbGorm *gorm.DB) error
}
