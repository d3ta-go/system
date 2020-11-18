package rdbms

import (
	"fmt"

	"github.com/d3ta-go/system/system/handler"
	"gorm.io/gorm"
)

// GormMigrationType represent GormMigrationType
type GormMigrationType string

const (
	GMTMigrate GormMigrationType = "MIGRATE"
	GMTSeed    GormMigrationType = "SEED"
)

// NewBaseGormMigrate create new BaseGormMigrate
func NewBaseGormMigrate(h *handler.Handler, dbConnName string) (*BaseGormMigrate, error) {
	var err error

	mig := new(BaseGormMigrate)
	mig.handler = h
	if err := mig.SetDBConnName(dbConnName); err != nil {
		return nil, err
	}
	if mig.migrationHistorySvc, err = NewGormMigrationHistoryService(h, mig.GetGorm()); err != nil {
		return nil, err
	}
	return mig, nil
}

// BaseGormMigrate represent BaseGormMigrate
type BaseGormMigrate struct {
	handler              *handler.Handler
	dbConnName           string
	dbGorm               *gorm.DB
	migrationHistorySvc  *GormMigrationHistoryService
	runMigrateItems      map[string][]IGormMigrateRunner
	rollBackMigrateItems map[string][]IGormMigrateRunner
}

// SetHandler set Handler
func (bgm *BaseGormMigrate) SetHandler(h *handler.Handler) {
	bgm.handler = h
}

// GetHandler get Handler
func (bgm *BaseGormMigrate) GetHandler() *handler.Handler {
	return bgm.handler
}

// SetDBConnName set dbConnName
func (bgm *BaseGormMigrate) SetDBConnName(dbConnName string) error {
	bgm.dbConnName = dbConnName
	if dbConnName != "" {
		dbGorm, err := bgm.handler.GetGormDB(bgm.dbConnName)
		if err != nil {
			return err
		}
		bgm.SetGorm(dbGorm)
	}
	return nil
}

// GetDBConnName get dbConnName
func (bgm *BaseGormMigrate) GetDBConnName() string {
	return bgm.dbConnName
}

// SetGorm set dbGorm
func (bgm *BaseGormMigrate) SetGorm(db *gorm.DB) {
	bgm.dbGorm = db
}

// GetGorm get dbGorm
func (bgm *BaseGormMigrate) GetGorm() *gorm.DB {
	return bgm.dbGorm
}

// GetRunMigrateItems get runMigrateItems
func (bgm *BaseGormMigrate) GetRunMigrateItems(mt GormMigrationType) []IGormMigrateRunner {
	val, ok := bgm.runMigrateItems[string(mt)]
	if !ok {
		var valnil []IGormMigrateRunner
		val = valnil
	}
	return val
}

// GetRollBackMigrateItems get rollBackMigrateItems
func (bgm *BaseGormMigrate) GetRollBackMigrateItems(mt GormMigrationType) []IGormMigrateRunner {
	val, ok := bgm.rollBackMigrateItems[string(mt)]
	if !ok {
		var valnil []IGormMigrateRunner
		val = valnil
	}
	return val
}

// RunMigrates run Migration
func (bgm *BaseGormMigrate) RunMigrates(h *handler.Handler, dbConnName string, items ...IGormMigrateRunner) error {
	return bgm._run(h, dbConnName, GMTMigrate, items...)
}

// RollBackMigrates rollback Migrate
func (bgm *BaseGormMigrate) RollBackMigrates(h *handler.Handler, dbConnName string, items ...IGormMigrateRunner) error {
	return bgm._rollback(h, dbConnName, GMTMigrate, items...)
}

// RunSeeds run Seed
func (bgm *BaseGormMigrate) RunSeeds(h *handler.Handler, dbConnName string, items ...IGormMigrateRunner) error {
	return bgm._run(h, dbConnName, GMTSeed, items...)
}

// RollBackSeeds roolback Seed
func (bgm *BaseGormMigrate) RollBackSeeds(h *handler.Handler, dbConnName string, items ...IGormMigrateRunner) error {
	return bgm._rollback(h, dbConnName, GMTSeed, items...)
}

func (bgm *BaseGormMigrate) _run(h *handler.Handler, dbConnName string, mt GormMigrationType, items ...IGormMigrateRunner) error {
	if h != nil {
		bgm.handler = h
	}
	if dbConnName != "" {
		bgm.dbConnName = dbConnName
		dbGorm, err := bgm.handler.GetGormDB(bgm.dbConnName)
		if err != nil {
			return err
		}
		bgm.dbGorm = dbGorm
	}

	if bgm.dbGorm != nil {
		if items != nil && len(items) > 0 {
			if bgm.runMigrateItems == nil {
				bgm.runMigrateItems = make(map[string][]IGormMigrateRunner)
			}
			var executedList []IGormMigrateRunner
			for _, ir := range items {
				if ir != nil {
					hasBeenExecuted, err := bgm.migrationHistorySvc.HasBeenExecuted(ir.GetID())
					if err != nil {
						return err
					}
					if hasBeenExecuted == false {
						if err := ir.Run(bgm.handler, bgm.dbGorm); err != nil {
							fmt.Printf("Error while running %s.Run: %s", ir.GetID(), err.Error())
							if err := bgm.migrationHistorySvc.SaveRunExecution(ir.GetID(), string(mt), err.Error()); err != nil {
								return err
							}
							return err
						}
						executedList = append(executedList, ir)
						if err := bgm.migrationHistorySvc.SaveRunExecution(ir.GetID(), string(mt), ""); err != nil {
							return err
						}
					}
				}
			}
			bgm.runMigrateItems[string(mt)] = executedList
		}
	}

	return nil
}

func (bgm *BaseGormMigrate) _rollback(h *handler.Handler, dbConnName string, mt GormMigrationType, items ...IGormMigrateRunner) error {
	if h != nil {
		bgm.handler = h
	}
	if dbConnName != "" {
		bgm.dbConnName = dbConnName
		dbGorm, err := bgm.handler.GetGormDB(bgm.dbConnName)
		if err != nil {
			return err
		}
		bgm.dbGorm = dbGorm
	}

	if bgm.dbGorm != nil {
		if items != nil && len(items) > 0 {
			if bgm.rollBackMigrateItems == nil {
				bgm.rollBackMigrateItems = make(map[string][]IGormMigrateRunner)
			}
			var executedList []IGormMigrateRunner
			for _, ir := range items {
				if ir != nil {
					hasBeenExecuted, err := bgm.migrationHistorySvc.HasBeenExecuted(ir.GetID())
					if err != nil {
						return err
					}
					if hasBeenExecuted {
						if err := ir.RollBack(bgm.handler, bgm.dbGorm); err != nil {
							fmt.Printf("Error while running %s.RollBack: %s", ir.GetID(), err.Error())
							if err := bgm.migrationHistorySvc.SaveRollBackExecution(ir.GetID(), err.Error()); err != nil {
								return err
							}
							return err
						}
						executedList = append(executedList, ir)
						if err := bgm.migrationHistorySvc.SaveRollBackExecution(ir.GetID(), ""); err != nil {
							return err
						}
					}
				}
			}
			bgm.rollBackMigrateItems[string(mt)] = executedList
		}
	}

	return nil
}
