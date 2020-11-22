package rdbms

import (
	"fmt"

	"github.com/d3ta-go/system/system/handler"
	"gorm.io/gorm"
)

// GormMigrationType represent GormMigrationType
type GormMigrationType string

const (
	// GMTMigrate represent GORM Migration Type - MIGRATE (ddl)
	GMTMigrate GormMigrationType = "MIGRATE"
	// GMTSeed represent GORM Migration TYpe - SEED (dml)
	GMTSeed GormMigrationType = "SEED"
)

// NewBaseGormMigrator create new BaseGormMigrator
func NewBaseGormMigrator(h *handler.Handler, dbConnName string) (*BaseGormMigrator, error) {
	var err error

	mig := new(BaseGormMigrator)
	mig.handler = h
	if err := mig.SetDBConnName(dbConnName); err != nil {
		return nil, err
	}
	if mig.migrationHistorySvc, err = NewGormMigrationHistoryService(h, mig.GetGorm()); err != nil {
		return nil, err
	}
	return mig, nil
}

// BaseGormMigrator represent BaseGormMigrator
type BaseGormMigrator struct {
	handler               *handler.Handler
	dbConnName            string
	dbGorm                *gorm.DB
	migrationHistorySvc   *GormMigrationHistoryService
	runMigratorItems      map[string][]IGormMigratorRunner
	rollBackMigratorItems map[string][]IGormMigratorRunner
}

// SetHandler set Handler
func (bgm *BaseGormMigrator) SetHandler(h *handler.Handler) {
	bgm.handler = h
}

// GetHandler get Handler
func (bgm *BaseGormMigrator) GetHandler() *handler.Handler {
	return bgm.handler
}

// SetDBConnName set dbConnName
func (bgm *BaseGormMigrator) SetDBConnName(dbConnName string) error {
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
func (bgm *BaseGormMigrator) GetDBConnName() string {
	return bgm.dbConnName
}

// SetGorm set dbGorm
func (bgm *BaseGormMigrator) SetGorm(db *gorm.DB) {
	bgm.dbGorm = db
}

// GetGorm get dbGorm
func (bgm *BaseGormMigrator) GetGorm() *gorm.DB {
	return bgm.dbGorm
}

// GetRunMigratorItems get runMigratorItems
func (bgm *BaseGormMigrator) GetRunMigratorItems(mt GormMigrationType) []IGormMigratorRunner {
	val, ok := bgm.runMigratorItems[string(mt)]
	if !ok {
		var valnil []IGormMigratorRunner
		val = valnil
	}
	return val
}

// GetRollBackMigratorItems get rollBackMigratorItems
func (bgm *BaseGormMigrator) GetRollBackMigratorItems(mt GormMigrationType) []IGormMigratorRunner {
	val, ok := bgm.rollBackMigratorItems[string(mt)]
	if !ok {
		var valnil []IGormMigratorRunner
		val = valnil
	}
	return val
}

// RunMigrates run Migration
func (bgm *BaseGormMigrator) RunMigrates(h *handler.Handler, dbConnName string, items ...IGormMigratorRunner) error {
	return bgm._run(h, dbConnName, GMTMigrate, items...)
}

// RollBackMigrates rollback Migrate
func (bgm *BaseGormMigrator) RollBackMigrates(h *handler.Handler, dbConnName string, items ...IGormMigratorRunner) error {
	return bgm._rollback(h, dbConnName, GMTMigrate, items...)
}

// RunSeeds run Seed
func (bgm *BaseGormMigrator) RunSeeds(h *handler.Handler, dbConnName string, items ...IGormMigratorRunner) error {
	return bgm._run(h, dbConnName, GMTSeed, items...)
}

// RollBackSeeds roolback Seed
func (bgm *BaseGormMigrator) RollBackSeeds(h *handler.Handler, dbConnName string, items ...IGormMigratorRunner) error {
	return bgm._rollback(h, dbConnName, GMTSeed, items...)
}

func (bgm *BaseGormMigrator) _run(h *handler.Handler, dbConnName string, mt GormMigrationType, items ...IGormMigratorRunner) error {
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
			if bgm.runMigratorItems == nil {
				bgm.runMigratorItems = make(map[string][]IGormMigratorRunner)
			}
			var executedList []IGormMigratorRunner
			for _, ir := range items {
				if ir != nil {
					hasBeenExecuted, err := bgm.migrationHistorySvc.HasBeenExecuted(ir.GetID())
					if err != nil {
						return err
					}
					if hasBeenExecuted == false {
						if err := ir.Run(bgm.handler, bgm.dbGorm); err != nil {
							fmt.Printf("Error while running %s.Run: %s\n", ir.GetID(), err.Error())
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
			bgm.runMigratorItems[string(mt)] = executedList
		}
	}

	return nil
}

func (bgm *BaseGormMigrator) _rollback(h *handler.Handler, dbConnName string, mt GormMigrationType, items ...IGormMigratorRunner) error {
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
			if bgm.rollBackMigratorItems == nil {
				bgm.rollBackMigratorItems = make(map[string][]IGormMigratorRunner)
			}
			var executedList []IGormMigratorRunner
			for _, ir := range items {
				if ir != nil {
					hasBeenExecuted, err := bgm.migrationHistorySvc.HasBeenExecuted(ir.GetID())
					if err != nil {
						return err
					}
					if hasBeenExecuted {
						if err := ir.RollBack(bgm.handler, bgm.dbGorm); err != nil {
							fmt.Printf("Error while running %s.RollBack: %s\n", ir.GetID(), err.Error())
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
			bgm.rollBackMigratorItems[string(mt)] = executedList
		}
	}

	return nil
}
