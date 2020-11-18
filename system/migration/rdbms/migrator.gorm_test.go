package rdbms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBaseGormMigrator(t *testing.T) (*BaseGormMigrator, error) {
	h, err := newHandler(t)
	if err != nil {
		return nil, err
	}

	m, err := NewBaseGormMigrator(h, "db-identity")
	if err != nil {
		return nil, err
	}

	return m, nil
}

func TestBaseGormMigrator_SetHandler(t *testing.T) {
	bgm := BaseGormMigrator{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
}

func TestBaseGormMigrator_GetHandler(t *testing.T) {
	bgm := BaseGormMigrator{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	assert.Equal(t, h, bgm.GetHandler())
}

func TestBaseGormMigrator_SetDBConnName(t *testing.T) {
	bgm := BaseGormMigrator{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	bgm.SetDBConnName("db-identity")
}

func TestBaseGormMigrator_GetDBConnName(t *testing.T) {
	bgm := BaseGormMigrator{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	bgm.SetDBConnName("db-identity")
	assert.Equal(t, "db-identity", bgm.GetDBConnName())
}

func TestBaseGormMigrator_SetGorm(t *testing.T) {
	bgm := BaseGormMigrator{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	bgm.SetDBConnName("db-identity")
	bgm.SetGorm(bgm.GetGorm())
}

func TestBaseGormMigrator_GetGorm(t *testing.T) {
	bgm := BaseGormMigrator{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	bgm.SetDBConnName("db-identity")

	gorm, err := h.GetGormDB("db-identity")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, gorm, bgm.GetGorm())
}

func TestBaseGormMigrator_GetRunMigratorItems(t *testing.T) {
	bgm := BaseGormMigrator{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	bgm.SetDBConnName("db-identity")

	var valnil []IGormMigratorRunner
	assert.Equal(t, valnil, bgm.GetRunMigratorItems(GMTMigrate))
}

func TestBaseGormMigrator_GetRollBackMigratorItems(t *testing.T) {
	bgm := BaseGormMigrator{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	bgm.SetDBConnName("db-identity")

	var valnil []IGormMigratorRunner
	assert.Equal(t, valnil, bgm.GetRollBackMigratorItems(GMTMigrate))
}

func TestBaseGormMigrator_RunMigrates(t *testing.T) {
	bgm, err := newBaseGormMigrator(t)
	if err != nil {
		t.Errorf("Error: newBaseGormMigrator(), [%s]", err.Error())
		return
	}
	if bgm != nil {
		newDummyMigrate, err := NewDummyMigrateRunner()
		if err != nil {
			t.Error(err)
		}
		assert.NoError(t, bgm.RunMigrates(bgm.GetHandler(), "db-identity", newDummyMigrate))
	}
}

func TestBaseGormMigrator_RunSeeds(t *testing.T) {
	bgm, err := newBaseGormMigrator(t)
	if err != nil {
		t.Errorf("Error: newBaseGormMigrator(), [%s]", err.Error())
		return
	}
	if bgm != nil {
		newDummySeed, err := NewDummySeedRunner()
		if err != nil {
			t.Error(err)
		}
		assert.NoError(t, bgm.RunSeeds(bgm.GetHandler(), "db-identity", newDummySeed))
	}
}

func TestBaseGormMigrator_RollBackSeeds(t *testing.T) {
	bgm, err := newBaseGormMigrator(t)
	if err != nil {
		t.Errorf("Error: newBaseGormMigrator(), [%s]", err.Error())
		return
	}
	if bgm != nil {
		newDummySeed, err := NewDummySeedRunner()
		if err != nil {
			t.Error(err)
		}
		assert.NoError(t, bgm.RollBackSeeds(bgm.GetHandler(), "db-identity", newDummySeed))
	}
}

func TestBaseGormMigrator_RollBackMigrates(t *testing.T) {
	bgm, err := newBaseGormMigrator(t)
	if err != nil {
		t.Errorf("Error: newBaseGormMigrator(), [%s]", err.Error())
		return
	}
	if bgm != nil {
		newDummyMigrate, err := NewDummyMigrateRunner()
		if err != nil {
			t.Error(err)
		}
		assert.NoError(t, bgm.RollBackMigrates(bgm.GetHandler(), "db-identity", newDummyMigrate))
	}
}
