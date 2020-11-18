package rdbms

import (
	"errors"
	"fmt"
	"time"

	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/utils"
	"gorm.io/gorm"
)

// GormMigrationHistoryEntity represent GormMigrationHistoryEntity
type GormMigrationHistoryEntity struct {
	gorm.Model
	ScriptName   string `gorm:"unique;size:500"`
	ScriptType   string `gorm:"size:50"`
	RunFrom      string `gorm:"size:200"`
	RunNote      string `gorm:"size:4000"`
	RollBackFrom string `gorm:"size:200"`
	RollBackNote string `gorm:"size:4000"`
}

// TableName get real database table name
func (t *GormMigrationHistoryEntity) TableName() string {
	return "_migration_histories"
}

// NewGormMigrationHistoryService create new GormMigrationHistoryService
func NewGormMigrationHistoryService(h *handler.Handler, db *gorm.DB) (*GormMigrationHistoryService, error) {
	if h == nil {
		return nil, fmt.Errorf("Handler should be not nil")
	}
	if db == nil {
		return nil, fmt.Errorf("DB should be not nill")
	}

	svc := new(GormMigrationHistoryService)
	svc.handler = h
	svc.dbGorm = db

	if err := svc.dbGorm.AutoMigrate(&GormMigrationHistoryEntity{}); err != nil {
		return nil, err
	}

	return svc, nil
}

// GormMigrationHistoryService represent GormMigrationHistoryService
type GormMigrationHistoryService struct {
	handler *handler.Handler
	dbGorm  *gorm.DB
}

// FindByScriptName find by ScriptName
func (s *GormMigrationHistoryService) FindByScriptName(scriptName string) (exist bool, ett *GormMigrationHistoryEntity, err error) {
	var ettMigHis GormMigrationHistoryEntity

	result := s.dbGorm.Unscoped().Where(&GormMigrationHistoryEntity{ScriptName: scriptName}).First(&ettMigHis)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil, result.Error
		}
	}

	return result.RowsAffected > 0, &ettMigHis, nil
}

// HasBeenExecuted has been executed (run)
func (s *GormMigrationHistoryService) HasBeenExecuted(scriptName string) (bool, error) {
	exist, _, err := s.FindByScriptName(scriptName)
	if err != nil {
		return false, err
	}
	return exist, nil
}

// SaveRunExecution save run execution
func (s *GormMigrationHistoryService) SaveRunExecution(scriptName string, scriptType string, note string) error {
	ettMigHis := GormMigrationHistoryEntity{
		ScriptName: scriptName,
		ScriptType: scriptType,
		RunNote:    note,
		RunFrom:    fmt.Sprintf("%s@%s", utils.GetHostName(), utils.GetCurrentIP()),
	}
	result := s.dbGorm.Create(&ettMigHis)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SaveRollBackExecution save rollback execution
func (s *GormMigrationHistoryService) SaveRollBackExecution(scriptName string, note string) error {
	exist, ettMigHis, err := s.FindByScriptName(scriptName)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("Migration script %s not executed before", scriptName)
	}

	ettMigHis.ScriptName = fmt.Sprintf("%s.roolback@%s", ettMigHis.ScriptName, time.Now().Format(time.RFC3339Nano))
	ettMigHis.RollBackNote = note
	ettMigHis.RollBackFrom = fmt.Sprintf("%s@%s", utils.GetHostName(), utils.GetCurrentIP())

	result := s.dbGorm.Save(&ettMigHis)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
