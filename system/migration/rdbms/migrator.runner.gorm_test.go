package rdbms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseGormMigratorRunner_SetID(t *testing.T) {
	bgmr := BaseGormMigratorRunner{}
	bgmr.SetID("unique-id")
}

func TestBaseGormMigratorRunner_GetID(t *testing.T) {
	bgmr := BaseGormMigratorRunner{}
	bgmr.SetID("unique-id")
	assert.Equal(t, "unique-id", bgmr.GetID())
}

func TestBaseGormMigratorRunner_SetHandler(t *testing.T) {
	bgmr := BaseGormMigratorRunner{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgmr.SetHandler(h)
}

func TestBaseGormMigratorRunner_GetHandler(t *testing.T) {
	bgmr := BaseGormMigratorRunner{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgmr.SetHandler(h)
	assert.Equal(t, h, bgmr.GetHandler())
}

func TestBaseGormMigratorRunner_SetDBConnName(t *testing.T) {
	bgmr := BaseGormMigratorRunner{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgmr.SetHandler(h)
	bgmr.SetDBConnName("db-identity")
}

func TestBaseGormMigratorRunner_GetDBConnName(t *testing.T) {
	bgmr := BaseGormMigratorRunner{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgmr.SetHandler(h)
	bgmr.SetDBConnName("db-identity")
	assert.Equal(t, "db-identity", bgmr.GetDBConnName())
}

func TestBaseGormMigratorRunner_SetGorm(t *testing.T) {
	bgmr := BaseGormMigratorRunner{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgmr.SetHandler(h)
	bgmr.SetDBConnName("db-identity")
	bgmr.SetGorm(bgmr.GetGorm())
}

func TestBaseGormMigratorRunner_GetGorm(t *testing.T) {
	bgmr := BaseGormMigratorRunner{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgmr.SetHandler(h)
	bgmr.SetDBConnName("db-identity")

	gorm, err := h.GetGormDB("db-identity")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, gorm, bgmr.GetGorm())
}
