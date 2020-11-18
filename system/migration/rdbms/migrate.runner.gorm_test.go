package rdbms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseGormMigrateRunner_SetID(t *testing.T) {
	bgmr := BaseGormMigrateRunner{}
	bgmr.SetID("unique-id")
}

func TestBaseGormMigrateRunner_GetID(t *testing.T) {
	bgmr := BaseGormMigrateRunner{}
	bgmr.SetID("unique-id")
	assert.Equal(t, "unique-id", bgmr.GetID())
}

func TestBaseGormMigrateRunner_SetHandler(t *testing.T) {
	bgmr := BaseGormMigrateRunner{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgmr.SetHandler(h)
}

func TestBaseGormMigrateRunner_GetHandler(t *testing.T) {
	bgmr := BaseGormMigrateRunner{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgmr.SetHandler(h)
	assert.Equal(t, h, bgmr.GetHandler())
}

func TestBaseGormMigrateRunner_SetDBConnName(t *testing.T) {
	bgmr := BaseGormMigrateRunner{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgmr.SetHandler(h)
	bgmr.SetDBConnName("db-identity")
}

func TestBaseGormMigrateRunner_GetDBConnName(t *testing.T) {
	bgmr := BaseGormMigrateRunner{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgmr.SetHandler(h)
	bgmr.SetDBConnName("db-identity")
	assert.Equal(t, "db-identity", bgmr.GetDBConnName())
}

func TestBaseGormMigrateRunner_SetGorm(t *testing.T) {
	bgmr := BaseGormMigrateRunner{}

	h, err := newHandler(t)
	if err != nil {
		t.Error(err)
	}
	bgmr.SetHandler(h)
	bgmr.SetDBConnName("db-identity")
	bgmr.SetGorm(bgmr.GetGorm())
}

func TestBaseGormMigrateRunner_GetGorm(t *testing.T) {
	bgmr := BaseGormMigrateRunner{}

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
