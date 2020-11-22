package elasticsearch

import (
	"fmt"
	"io"

	"github.com/d3ta-go/system/system/indexer/adapter"

	es6 "github.com/elastic/go-elasticsearch/v6"
	es7 "github.com/elastic/go-elasticsearch/v7"
	es8 "github.com/elastic/go-elasticsearch/v8"
)

// ESVersion represent Elastic Version Type
type ESVersion string

const (
	// ESVersion6 represent Elastic Search Version 6
	ESVersion6 ESVersion = "6"
	// ESVersion7 represent Elastic Search Version 7
	ESVersion7 ESVersion = "7"
	// ESVersion8 represent Elastic Search Version 8
	ESVersion8 ESVersion = "8"
)

// NewIndexer new Elastic Search Indexer
func NewIndexer(version ESVersion, cfg interface{}) (adapter.IIndexerEngine, error) {
	var err error

	idx := &Indexer{
		esVersion: version,
	}

	cfgType := fmt.Sprintf("%T", cfg)
	if cfgType != "elasticsearch.Config" {
		return nil, fmt.Errorf("Invalid Configuration Type (should be: `elasticsearch.Config`)")
	}

	switch version {
	case ESVersion6:
		if idx.engine, err = NewIndexerES6(cfg.(es6.Config)); err != nil {
			return nil, err
		}
	case ESVersion7:
		if idx.engine, err = NewIndexerES7(cfg.(es7.Config)); err != nil {
			return nil, err
		}
	case ESVersion8:
		if idx.engine, err = NewIndexerES8(cfg.(es8.Config)); err != nil {
			return nil, err
		}
	}

	return idx, nil
}

// Indexer type
type Indexer struct {
	esVersion ESVersion
	engine    adapter.IIndexerEngine
}

// GetEngine get Indexer Engine
func (i *Indexer) GetEngine() interface{} {
	return i.engine.GetEngine()
}

// Search is a function to Search
func (i *Indexer) Search(query io.Reader, prettify bool) ([]byte, error) {
	return i.engine.Search(query, prettify)
}

// SearchIndexDoc is a function to Search Index Doc
func (i *Indexer) SearchIndexDoc(index string, query io.Reader, size int, prettify bool) ([]byte, error) {
	return i.engine.SearchIndexDoc(index, query, size, prettify)
}

// IndexExist is a function to check Index is Exist
func (i *Indexer) IndexExist(indexs []string) (bool, error) {
	return i.engine.IndexExist(indexs)
}

// CreateIndex is a function to Create Index
func (i *Indexer) CreateIndex(index string, mapping io.Reader) error {
	return i.engine.CreateIndex(index, mapping)
}

// DropIndex is a function to Drop Index
func (i *Indexer) DropIndex(indexs []string) error {
	return i.engine.DropIndex(indexs)
}

// DocExist is a function to check Doc is Exist
func (i *Indexer) DocExist(index string, id string) (bool, error) {
	return i.engine.DocExist(index, id)
}

// CreateDoc is a function to Create Doc
func (i *Indexer) CreateDoc(index string, id string, body io.Reader) error {

	exist, err := i.DocExist(index, id)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("Index Document already exist: %s[id=%s]", index, id)
	}
	return i.engine.CreateDoc(index, id, body)
}

// GetDoc is a function to Get Doc
func (i *Indexer) GetDoc(index string, id string) ([]byte, error) {
	return i.engine.GetDoc(index, id)
}

// UpdateDoc is a function to Update Doc
func (i *Indexer) UpdateDoc(index string, id string, body io.Reader) error {

	exist, err := i.DocExist(index, id)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("Index Document does not exist: %s[id=%s]", index, id)
	}
	return i.engine.UpdateDoc(index, id, body)
}

// DeleteDoc is a function to Delete Doc
func (i *Indexer) DeleteDoc(index string, id string) error {
	exist, err := i.DocExist(index, id)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("Index Document does not exist: %s[id=%s]", index, id)
	}
	return i.engine.DeleteDoc(index, id)
}
