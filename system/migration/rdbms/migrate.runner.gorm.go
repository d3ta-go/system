package rdbms

import (
	"github.com/d3ta-go/system/system/handler"
	"gorm.io/gorm"
)

// BaseGormMigrateRunner represent BaseGormMigrateRunner
type BaseGormMigrateRunner struct {
	id         string
	handler    *handler.Handler
	dbConnName string
	dbGorm     *gorm.DB
}

// SetID set ID
func (bmr *BaseGormMigrateRunner) SetID(id string) {
	bmr.id = id
}

// GetID get ID
func (bmr *BaseGormMigrateRunner) GetID() string {
	return bmr.id
}

// SetHandler set Handler
func (bmr *BaseGormMigrateRunner) SetHandler(h *handler.Handler) {
	bmr.handler = h
}

// GetHandler get Handler
func (bmr *BaseGormMigrateRunner) GetHandler() *handler.Handler {
	return bmr.handler
}

// SetDBConnName set dbConnName
func (bmr *BaseGormMigrateRunner) SetDBConnName(dbConnName string) error {
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
func (bmr *BaseGormMigrateRunner) GetDBConnName() string {
	return bmr.dbConnName
}

// SetGorm set dbGorm
func (bmr *BaseGormMigrateRunner) SetGorm(db *gorm.DB) {
	bmr.dbGorm = db
}

// GetGorm get dbGorm
func (bmr *BaseGormMigrateRunner) GetGorm() *gorm.DB {
	return bmr.dbGorm
}
