package initialize

import (
	"reflect"

	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/indexer"
	"github.com/d3ta-go/system/system/indexer/adapter"
)

// OpenAllIndexerConnection open all indexer connection
func OpenAllIndexerConnection(h *handler.Handler) error {
	cfg, err := h.GetDefaultConfig()
	if err != nil {
		return err
	}

	if cfg != nil {
		idxs := cfg.Indexers
		e := reflect.ValueOf(&idxs).Elem()
		for i := 0; i < e.NumField(); i++ {
			idxConfig := e.Field(i).Interface()
			if idxConfig != nil {
				err := OpenIndexerConnection(idxConfig.(config.Indexer), h)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// OpenIndexerConnection open indexer connection
func OpenIndexerConnection(config config.Indexer, h *handler.Handler) error {
	if h != nil {
		options := config.Configurations
		ie, it, err := indexer.NewIndexerEngine(indexer.TheIndexerType(config.Driver), adapter.IEOptions{
			Engine:  adapter.IEType(config.Engine),
			Version: config.Version,
			Options: options,
		})
		if err != nil {
			return err
		}

		idx, err := indexer.NewIndexer(it, ie)
		if err != nil {
			return err
		}
		h.SetIndexer(config.ConnectionName, idx)
	}

	return nil
}
