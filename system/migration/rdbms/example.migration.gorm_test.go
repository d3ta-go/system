package rdbms

import (
	"fmt"

	"github.com/d3ta-go/system/system/handler"
	"gorm.io/gorm"
)

type DummyTable struct {
	gorm.Model
	Dummy string `gorm:"index;size:255"`
}

// =============
// MigrateRunner
// =============
type DummyMigrateRunner struct {
	BaseGormMigrateRunner
}

func NewDummyMigrateRunner() (IGormMigrateRunner, error) {
	gmr := new(DummyMigrateRunner)
	gmr.SetID("DummyMigrateRunner")
	return gmr, nil
}

func (dmr *DummyMigrateRunner) GetID() string {
	return fmt.Sprintf("%T", dmr)
}

func (dmr *DummyMigrateRunner) Run(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.dbGorm = dbGorm
	}
	if dmr.dbGorm != nil {
		if err := dmr.dbGorm.AutoMigrate(&DummyTable{}); err != nil {
			return err
		}
	}
	return nil
}

func (dmr *DummyMigrateRunner) RollBack(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.dbGorm = dbGorm
	}
	if dmr.dbGorm != nil {
		if dmr.dbGorm.Migrator().HasTable(&DummyTable{}) {
			if err := dmr.dbGorm.Migrator().DropTable(&DummyTable{}); err != nil {
				return err
			}
		}
	}
	return nil
}

// =============
// SeedRunner
// =============
type DummySeedRunner struct {
	BaseGormMigrateRunner
}

func NewDummySeedRunner() (IGormMigrateRunner, error) {
	gms := new(DummySeedRunner)
	gms.SetID("DummySeedRunner")
	return gms, nil
}

func (dmr *DummySeedRunner) GetID() string {
	return fmt.Sprintf("%T", dmr)
}

func (dmr *DummySeedRunner) Run(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.dbGorm = dbGorm
	}
	if dmr.dbGorm != nil {
		if dmr.dbGorm.Migrator().HasTable(&DummyTable{}) {
			result := dmr.dbGorm.Create(&DummyTable{Dummy: "Dummy Data"})
			if result.Error != nil {
				return result.Error
			}
		} else {
			return fmt.Errorf("table not found")
		}
	}
	return nil
}

func (dmr *DummySeedRunner) RollBack(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.dbGorm = dbGorm
	}
	if dmr.dbGorm != nil {
		if dmr.dbGorm.Migrator().HasTable(&DummyTable{}) {
			result := dmr.dbGorm.Unscoped().Where(&DummyTable{Dummy: "Dummy Data"}).Delete(&DummyTable{})
			if result.Error != nil {
				return result.Error
			}
		} else {
			return fmt.Errorf("table not found")
		}
	}
	return nil
}
