package indexer

import (
	"fmt"

	"github.com/d3ta-go/system/system/indexer/adapter"
	ieES "github.com/d3ta-go/system/system/indexer/adapter/elasticsearch"
)

// NewIndexerEngine new IndexerEngine
func NewIndexerEngine(TheIndexerType TheIndexerType, ieOptions adapter.IEOptions) (adapter.IIndexerEngine, TheIndexerType, error) {
	tOpt := fmt.Sprintf("%T", ieOptions.Options)

	switch ieOptions.Engine {
	case adapter.IEElasticSearch:
		if tOpt != "map[string]interface {}" {
			return nil, "", fmt.Errorf("Invalid ElasticSearch options (should be: `map[string]interface {}`) [actual: `%s`]", tOpt)
		}

		var ie adapter.IIndexerEngine
		var err error

		switch ieOptions.Version {
		case "6":
			cfg := ieES.ConfigParserES6(ieOptions.Options)
			ie, err = ieES.NewIndexer(ieES.ESVersion6, cfg)
		case "7":
			cfg := ieES.ConfigParserES7(ieOptions.Options)
			ie, err = ieES.NewIndexer(ieES.ESVersion7, cfg)
		case "8":
			cfg := ieES.ConfigParserES8(ieOptions.Options)
			ie, err = ieES.NewIndexer(ieES.ESVersion8, cfg)
		default:
			err = fmt.Errorf("Invalid adapter.IEElasticSearch Version: %s", ieOptions.Version)
		}

		if err != nil {
			return nil, "", err
		}
		return ie, TheIndexerType, nil

	default:
		return nil, "", fmt.Errorf("Invalid Indexer Engine")
	}
}
