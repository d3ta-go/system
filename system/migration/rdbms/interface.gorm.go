package rdbms

import (
	"github.com/d3ta-go/system/system/handler"
	"gorm.io/gorm"
)

// IGormMigratorRunner represent GormMigratorRunner Interface
type IGormMigratorRunner interface {
	GetID() string
	Run(h *handler.Handler, dbGorm *gorm.DB) error
	RollBack(h *handler.Handler, dbGorm *gorm.DB) error
}
