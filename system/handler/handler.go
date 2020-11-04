package handler

import (
	"fmt"

	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"

	"github.com/casbin/casbin/v2"
	"github.com/d3ta-go/system/system/cacher"
	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/indexer"
)

// NewHandler new Handler
func NewHandler() (*Handler, error) {
	h := new(Handler)

	h.dbGorms = make(map[string]*gorm.DB)

	return h, nil
}

// Handler represent Handler
type Handler struct {
	defaultConfig   *config.Config
	dbGorms         map[string]*gorm.DB
	casbinEnforcers map[string]*casbin.Enforcer
	cachers         map[string]*cacher.Cacher
	indexers        map[string]*indexer.Indexer
}

// SetDefaultConfig set DefaultConfig
func (h *Handler) SetDefaultConfig(config *config.Config) {
	h.defaultConfig = config
}

// GetDefaultConfig get DefaultConfig
func (h *Handler) GetDefaultConfig() (*config.Config, error) {
	if h.defaultConfig == nil {
		return nil, fmt.Errorf("ERROR: [%s]", "Default Configuration Does Not Exist")
	}
	return h.defaultConfig, nil
}

// SetGormDB set GORM database connection by connection name
func (h *Handler) SetGormDB(conName string, dbCon *gorm.DB) {
	if h.dbGorms == nil {
		h.dbGorms = make(map[string]*gorm.DB)
	}
	h.dbGorms[conName] = dbCon
}

// GetGormDB get GORM database connection by connection name
func (h *Handler) GetGormDB(conName string) (*gorm.DB, error) {
	db, exist := h.dbGorms[conName]
	if !exist {
		err := fmt.Errorf("DB Connection Name '%s' Not Found", conName)
		return nil, err
	}
	return db, nil
}

// GetGormDBs get Gorm Databases
func (h *Handler) GetGormDBs() map[string]*gorm.DB {
	return h.dbGorms
}

// SetCasbinEnforcer set CasbinEnforcer
func (h *Handler) SetCasbinEnforcer(ceName string, ce *casbin.Enforcer) {
	if h.casbinEnforcers == nil {
		h.casbinEnforcers = make(map[string]*casbin.Enforcer)
	}
	h.casbinEnforcers[ceName] = ce
}

// GetCasbinEnforcer get CasbinEnforcer
func (h *Handler) GetCasbinEnforcer(ceName string) (*casbin.Enforcer, error) {
	ce, exist := h.casbinEnforcers[ceName]
	if !exist {
		err := fmt.Errorf("Casbin Enforcer Name '%s' Not Found", ceName)
		return nil, err
	}
	return ce, nil
}

// SetCacher set Cacher
func (h *Handler) SetCacher(cName string, c *cacher.Cacher) {
	if h.cachers == nil {
		h.cachers = make(map[string]*cacher.Cacher)
	}
	h.cachers[cName] = c
}

// GetCacher get Cacher
func (h *Handler) GetCacher(cName string) (*cacher.Cacher, error) {
	c, exist := h.cachers[cName]
	if !exist {
		err := fmt.Errorf("Cacher Name '%s' Not Found", cName)
		return nil, err
	}
	return c, nil
}

// GetCachers get Cachers
func (h *Handler) GetCachers() map[string]*cacher.Cacher {
	return h.cachers
}

// SetIndexer set Indexer
func (h *Handler) SetIndexer(idxName string, idx *indexer.Indexer) {
	if h.indexers == nil {
		h.indexers = make(map[string]*indexer.Indexer)
	}
	h.indexers[idxName] = idx
}

// GetIndexer get Indexer
func (h *Handler) GetIndexer(idxName string) (*indexer.Indexer, error) {
	idx, exist := h.indexers[idxName]
	if !exist {
		err := fmt.Errorf("Indexer Name '%s' Not Found", idxName)
		return nil, err
	}
	return idx, nil
}

// GetIndexers get Indexers
func (h *Handler) GetIndexers() map[string]*indexer.Indexer {
	return h.indexers
}
