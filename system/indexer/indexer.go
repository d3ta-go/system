package indexer

import (
	"fmt"
	"io"
	"strings"

	"github.com/d3ta-go/system/system/indexer/adapter"
)

// TheIndexerType represent TheIndexerType
type TheIndexerType string

const (
	// ElasticSearchIndexer represent ElasticSearch Indexer
	ElasticSearchIndexer TheIndexerType = "elasticsearch"
)

// NewIndexer create new Indexer
func NewIndexer(TheIndexerType TheIndexerType, indexerEngine adapter.IIndexerEngine) (*Indexer, error) {
	if indexerEngine == nil {
		return nil, fmt.Errorf("Invalid indexerEngine value")
	}

	idx := Indexer{
		_type:         TheIndexerType,
		indexerEngine: indexerEngine,
	}

	// C4 prefix key
	idx.Context = "defaultContext"
	idx.Container = "defaultContainer"
	idx.Component = "defaultComponent"
	// idx.Code = ...

	return &idx, nil
}

// Indexer type
type Indexer struct {
	_type         TheIndexerType
	indexerEngine adapter.IIndexerEngine

	Context   string
	Container string
	Component string
}

// GetEngine represent GetEngine
func (i *Indexer) GetEngine() interface{} {
	return i.indexerEngine.GetEngine()
}

// Search represent Search
func (i *Indexer) Search(query io.Reader, prettify bool) ([]byte, error) {
	return i.indexerEngine.Search(query, prettify)
}

// SearchIndexDoc represent SearchIndexDoc
func (i *Indexer) SearchIndexDoc(index string, query io.Reader, size int, prettify bool) ([]byte, error) {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.SearchIndexDoc(_index, query, size, prettify)
}

// IndexExist represent IndexExist
func (i *Indexer) IndexExist(indexs []string) (bool, error) {
	var _indexs []string
	for _, v := range indexs {
		_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), v)
		_indexs = append(_indexs, _index)
	}
	return i.indexerEngine.IndexExist(_indexs)
}

// CreateIndex represent CreateIndex
func (i *Indexer) CreateIndex(index string, mapping io.Reader) error {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.CreateIndex(_index, mapping)
}

// DropIndex represent DropIndex
func (i *Indexer) DropIndex(indexs []string) error {
	var _indexs []string
	for _, v := range indexs {
		_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), v)
		_indexs = append(_indexs, _index)
	}
	return i.indexerEngine.DropIndex(_indexs)
}

// DocExist represent DocExist
func (i *Indexer) DocExist(index string, id string) (bool, error) {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.DocExist(_index, id)
}

// CreateDoc represent CreateDoc
func (i *Indexer) CreateDoc(index string, id string, body io.Reader) error {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.CreateDoc(_index, id, body)
}

// GetDoc represent GetDoc
func (i *Indexer) GetDoc(index string, id string) ([]byte, error) {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.GetDoc(_index, id)
}

// UpdateDoc represent UpdateDoc
func (i *Indexer) UpdateDoc(index string, id string, body io.Reader) error {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.UpdateDoc(_index, id, body)
}

// DeleteDoc represent DeleteDoc
func (i *Indexer) DeleteDoc(index string, id string) error {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.DeleteDoc(_index, id)
}
