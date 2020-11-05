package initialize

import (
	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/indexer"
	"github.com/d3ta-go/system/system/indexer/adapter"
)

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
