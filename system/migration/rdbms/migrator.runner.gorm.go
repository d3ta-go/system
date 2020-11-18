package rdbms

import (
	"github.com/d3ta-go/system/system/handler"
	"gorm.io/gorm"
)

// BaseGormMigratorRunner represent BaseGormMigratorRunner
type BaseGormMigratorRunner struct {
	id         string
	handler    *handler.Handler
	dbConnName string
	dbGorm     *gorm.DB
}

// SetID set ID
func (bmr *BaseGormMigratorRunner) SetID(id string) {
	bmr.id = id
}

// GetID get ID
func (bmr *BaseGormMigratorRunner) GetID() string {
	return bmr.id
}

// SetHandler set Handler
func (bmr *BaseGormMigratorRunner) SetHandler(h *handler.Handler) {
	bmr.handler = h
}

// GetHandler get Handler
func (bmr *BaseGormMigratorRunner) GetHandler() *handler.Handler {
	return bmr.handler
}

// SetDBConnName set dbConnName
func (bmr *BaseGormMigratorRunner) SetDBConnName(dbConnName string) error {
	bmr.dbConnName = dbConnName
	if dbConnName != "" {
		dbGorm, err := bmr.handler.GetGormDB(bmr.dbConnName)
		if err != nil {
			return err
		}
		bmr.SetGorm(dbGorm)
	}
	return nil
}

// GetDBConnName get dbConnName
func (bmr *BaseGormMigratorRunner) GetDBConnName() string {
	return bmr.dbConnName
}

// SetGorm set dbGorm
func (bmr *BaseGormMigratorRunner) SetGorm(db *gorm.DB) {
	bmr.dbGorm = db
}

// GetGorm get dbGorm
func (bmr *BaseGormMigratorRunner) GetGorm() *gorm.DB {
	return bmr.dbGorm
}
