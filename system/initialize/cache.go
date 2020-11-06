package initialize

import (
	"fmt"
	"reflect"

	"github.com/d3ta-go/system/system/cacher"
	"github.com/d3ta-go/system/system/cacher/adapter"
	ceGoMacaron "github.com/d3ta-go/system/system/cacher/adapter/gomacaron"
	ceRedis "github.com/d3ta-go/system/system/cacher/adapter/redis"
	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
)

// OpenAllCacheConnection open all cache connection
func OpenAllCacheConnection(h *handler.Handler) error {
	cfg, err := h.GetDefaultConfig()
	if err != nil {
		return err
	}

	if cfg != nil {
		dbs := cfg.Caches
		e := reflect.ValueOf(&dbs).Elem()
		for i := 0; i < e.NumField(); i++ {
			cacheConfig := e.Field(i).Interface()
			if cacheConfig != nil {
				err := OpenCacheConnection(cacheConfig.(config.Cache), h)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// OpenCacheConnection open CacheConnection
func OpenCacheConnection(config config.Cache, h *handler.Handler) error {
	if h != nil {
		switch config.Engine {
		case "redis":
			options := ceRedis.ConfigParser(config.Configurations)

			ce, ct, err := cacher.NewCacherEngine(cacher.RedisCacher, adapter.CEOptions{
				Engine:  adapter.CERedis,
				Options: options,
			})
			if err != nil {
				return err
			}

			c, err := cacher.NewCacher(ct, ce)
			if err != nil {
				return err
			}
			h.SetCacher(config.ConnectionName, c)

		case "gomacaron":
			options := ceGoMacaron.ConfigParser(config.Configurations)

			ce, ct, err := cacher.NewCacherEngine(cacher.TheCacherType(config.Driver), adapter.CEOptions{
				Engine:  adapter.CEGoMacaron,
				Options: options,
			})
			if err != nil {
				return err
			}

			c, err := cacher.NewCacher(ct, ce)
			if err != nil {
				return err
			}
			h.SetCacher(config.ConnectionName, c)

		default:
			return fmt.Errorf("Invalid Cacher Engine")
		}
	}

	return nil
}
