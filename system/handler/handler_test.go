package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	h, err := NewHandler()
	if !assert.NoError(t, err, "Error while create new Handler: NewHandler") {
		return
	}

	// Config
	h.SetDefaultConfig(nil)
	cfg, err := h.GetDefaultConfig()
	if assert.Error(t, err, "Should be Error while getting config from Handler: h.GetDefaultConfig()") {
		assert.Nil(t, cfg)
	}

	// SpecificConfig
	type Me struct {
		One  string
		Two  string
		Tree string
	}

	h.SetSpecificConfig("try", Me{One: "Satu", Two: "Dua", Tree: "Tiga"})
	sCfg, err := h.GetSpecificConfig("try")
	if assert.NoError(t, err, "Should be No Error while getting config from Handler: h.GetSpecificConfig()") {
		assert.NotNil(t, sCfg)
		me, ok := sCfg.(Me)
		if !ok {
			t.Error("invalid sCfg")
		}
		assert.NotEmpty(t, me.One)
	}

	// Viper
	viper, err := h.GetViper("not-found")
	if assert.Error(t, err, "Should be Error while getting Viper from Handler: h.GetViper()") {
		assert.Nil(t, viper)
	}

	err2 := h.SetViper("nil-value", nil)
	if assert.Error(t, err2, "Should be Error while setting Viper from Handler: h.SetViper()") {
		assert.Nil(t, viper)
	}

	// Gorm DB
	dbCon, err := h.GetGormDB("not-found")
	if assert.Error(t, err, "Should be Error while getting GormDB from Handler: h.GetGormDB()") {
		assert.Nil(t, dbCon)
	}

	h.SetGormDB("nil-value", nil)
	dbCon2, err2 := h.GetGormDB("nil-value")
	if assert.NoError(t, err2, "Error while getting GormDB from Handler: h.GetGormDB()") {
		assert.Nil(t, dbCon2)
	}

	// Casbin Enforcer
	ce, err := h.GetCasbinEnforcer("not-found")
	if assert.Error(t, err, "Should be Error while getting Casbin Enforcer from Handler: h.GetCasbinEnforcer()") {
		assert.Nil(t, ce)
	}

	h.SetCasbinEnforcer("nil-value", nil)
	ce2, err2 := h.GetCasbinEnforcer("nil-value")
	if assert.NoError(t, err2, "Error while getting Casbin Enforcer from Handler: h.GetCasbinEnforcer()") {
		assert.Nil(t, ce2)
	}

	// Cacher
	c, err := h.GetCacher("not-found")
	if assert.Error(t, err, "Should be Error while getting Cacher from Handler: h.GetCacher()") {
		assert.Nil(t, c)
	}

	h.SetCacher("nil-value", nil)
	c2, err2 := h.GetCacher("nil-value")
	if assert.NoError(t, err2, "Error while getting Cacher from Handler: h.GetCacher()") {
		assert.Nil(t, c2)
	}

	// Indexer
	idx, err := h.GetIndexer("not-found")
	if assert.Error(t, err, "Should be Error while getting Indexer from Handler: h.GetIndexer()") {
		assert.Nil(t, idx)
	}

	h.SetIndexer("nil-value", nil)
	idx2, err2 := h.GetIndexer("nil-value")
	if assert.NoError(t, err2, "Error while getting Indexer from Handler: h.GetIndexer()") {
		assert.Nil(t, idx2)
	}
}
