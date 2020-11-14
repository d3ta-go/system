package handler

import (
	"fmt"

	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"

	"github.com/casbin/casbin/v2"
	"github.com/d3ta-go/system/system/cacher"
	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/indexer"
	"github.com/spf13/viper"
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
	sConfigs        map[string]interface{}
	vipers          map[string]*viper.Viper
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

// SetSpecificConfig set specific config by config name
func (h *Handler) SetSpecificConfig(cfgName string, cfg interface{}) {
	if h.sConfigs == nil {
		h.sConfigs = make(map[string]interface{})
	}
	h.sConfigs[cfgName] = cfg
}

// GetSpecificConfig get specific config by config name
func (h *Handler) GetSpecificConfig(cfgName string) (interface{}, error) {
	cfg, exist := h.sConfigs[cfgName]
	if !exist {
		err := fmt.Errorf("Config Name '%s' Not Found", cfgName)
		return nil, err
	}
	return cfg, nil
}

// SetViper set viper
func (h *Handler) SetViper(key string, v *viper.Viper) error {
	if v == nil {
		return fmt.Errorf("invalid viper")
	}
	if h.vipers == nil {
		h.vipers = make(map[string]*viper.Viper)
	}
	h.vipers[key] = v
	return nil
}

// GetViper get viper
func (h *Handler) GetViper(key string) (*viper.Viper, error) {
	v, exist := h.vipers[key]
	if !exist {
		err := fmt.Errorf("Viper Key '%s' Not Found", key)
		return nil, err
	}
	return v, nil
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
