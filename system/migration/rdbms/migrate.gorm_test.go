package rdbms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBaseGormMigrate(t *testing.T) (*BaseGormMigrate, error) {
	h, err := newHandler(t)
	if err != nil {
		return nil, err
	}

	m, err := NewBaseGormMigrate(h, "db-identity")
	if err != nil {
		return nil, err
	}

	return m, nil
}

func TestBaseGormMigrate_SetHandler(t *testing.T) {
	bgm := BaseGormMigrate{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
}

func TestBaseGormMigrate_GetHandler(t *testing.T) {
	bgm := BaseGormMigrate{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	assert.Equal(t, h, bgm.GetHandler())
}

func TestBaseGormMigrate_SetDBConnName(t *testing.T) {
	bgm := BaseGormMigrate{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	bgm.SetDBConnName("db-identity")
}

func TestBaseGormMigrate_GetDBConnName(t *testing.T) {
	bgm := BaseGormMigrate{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	bgm.SetDBConnName("db-identity")
	assert.Equal(t, "db-identity", bgm.GetDBConnName())
}

func TestBaseGormMigrate_SetGorm(t *testing.T) {
	bgm := BaseGormMigrate{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	bgm.SetDBConnName("db-identity")
	bgm.SetGorm(bgm.GetGorm())
}

func TestBaseGormMigrate_GetGorm(t *testing.T) {
	bgm := BaseGormMigrate{}

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

func TestBaseGormMigrate_GetRunMigrateItems(t *testing.T) {
	bgm := BaseGormMigrate{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	bgm.SetDBConnName("db-identity")

	var valnil []IGormMigrateRunner
	assert.Equal(t, valnil, bgm.GetRunMigrateItems(GMTMigrate))
}

func TestBaseGormMigrate_GetRollBackMigrateItems(t *testing.T) {
	bgm := BaseGormMigrate{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgm.SetHandler(h)
	bgm.SetDBConnName("db-identity")

	var valnil []IGormMigrateRunner
	assert.Equal(t, valnil, bgm.GetRollBackMigrateItems(GMTMigrate))
}

func TestBaseGormMigrate_RunMigrates(t *testing.T) {
	bgm, err := newBaseGormMigrate(t)
	if err != nil {
		t.Errorf("Error: newBaseGormMigrate(), [%s]", err.Error())
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

func TestBaseGormMigrate_RunSeeds(t *testing.T) {
	bgm, err := newBaseGormMigrate(t)
	if err != nil {
		t.Errorf("Error: newBaseGormMigrate(), [%s]", err.Error())
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

func TestBaseGormMigrate_RollBackSeeds(t *testing.T) {
	bgm, err := newBaseGormMigrate(t)
	if err != nil {
		t.Errorf("Error: newBaseGormMigrate(), [%s]", err.Error())
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

func TestBaseGormMigrate_RollBackMigrates(t *testing.T) {
	bgm, err := newBaseGormMigrate(t)
	if err != nil {
		t.Errorf("Error: newBaseGormMigrate(), [%s]", err.Error())
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
