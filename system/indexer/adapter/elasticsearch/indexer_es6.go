package elasticsearch

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/d3ta-go/system/system/indexer/adapter"
	es6 "github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
)

// NewIndexerES6 new Elastic Search 6 Indexer
func NewIndexerES6(cfg es6.Config) (adapter.IIndexerEngine, error) {
	var err error

	idx := &IndexerES6{
		esVersion: ESVersion6,
	}
	idx.ctx = context.Background()
	if idx.engine, err = es6.NewClient(cfg); err != nil {
		return nil, err
	}

	return idx, nil
}

// IndexerES6 type
type IndexerES6 struct {
	ctx       context.Context
	esVersion ESVersion
	engine    *es6.Client
}

// GetEngine get Indexer Engine
func (i *IndexerES6) GetEngine() interface{} {
	return i.engine
}

// Search is a function to Search
func (i *IndexerES6) Search(query io.Reader, prettify bool) ([]byte, error) {
	var res *esapi.Response
	var err error

	if prettify {
		res, err = i.engine.Search(
			i.engine.Search.WithContext(i.ctx),
			i.engine.Search.WithBody(query),
			i.engine.Search.WithPretty(),
		)
	} else {
		res, err = i.engine.Search(
			i.engine.Search.WithContext(i.ctx),
			i.engine.Search.WithBody(query),
		)
	}
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// SearchIndexDoc is a function to Search Index Doc
func (i *IndexerES6) SearchIndexDoc(index string, query io.Reader, size int, prettify bool) ([]byte, error) {
	var res *esapi.Response
	var err error

	if prettify {
		res, err = i.engine.Search(
			i.engine.Search.WithContext(i.ctx),
			i.engine.Search.WithIndex(index),
			i.engine.Search.WithBody(query),
			i.engine.Search.WithSize(size),
			i.engine.Search.WithPretty(),
		)

	} else {
		res, err = i.engine.Search(
			i.engine.Search.WithContext(i.ctx),
			i.engine.Search.WithIndex(index),
			i.engine.Search.WithBody(query),
			i.engine.Search.WithSize(size),
		)
	}
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// IndexExist is a function to check Index is Exist
func (i *IndexerES6) IndexExist(indexs []string) (bool, error) {
	res, err := i.engine.Indices.Exists(indexs)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	return res.StatusCode == 200, nil
}

// CreateIndex is a function to Create Index
func (i *IndexerES6) CreateIndex(index string, mapping io.Reader) error {
	res, err := i.engine.Indices.Create(index, i.engine.Indices.Create.WithBody(mapping))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

// DropIndex is a function to Drop Index
func (i *IndexerES6) DropIndex(indexs []string) error {
	res, err := i.engine.Indices.Delete(indexs)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

// DocExist is a function to check Doc is Exist
func (i *IndexerES6) DocExist(index string, id string) (bool, error) {

	res, err := i.engine.Exists(index, id)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	return res.StatusCode == 200, nil
}

// CreateDoc is a function to Create Doc
func (i *IndexerES6) CreateDoc(index string, id string, body io.Reader) error {

	res, err := i.engine.Create(index, id, body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

// GetDoc is a function to Get Doc
func (i *IndexerES6) GetDoc(index string, id string) ([]byte, error) {

	res, err := i.engine.Get(index, id)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// UpdateDoc is a function to Update Doc
func (i *IndexerES6) UpdateDoc(index string, id string, body io.Reader) error {

	res, err := i.engine.Update(index, id, body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

// DeleteDoc is a function to Delete Doc
func (i *IndexerES6) DeleteDoc(index string, id string) error {

	res, err := i.engine.Delete(index, id)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", string(body))
	}

	return nil
}
