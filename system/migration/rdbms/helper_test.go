package rdbms

import (
	"fmt"
	"testing"

	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func newConfig(t *testing.T) (*config.Config, *viper.Viper, error) {
	c, v, err := config.NewConfig("../../../conf")
	if err != nil {
		return nil, nil, err
	}
	if !c.CanRunTest() {
		panic(fmt.Sprintf("Cannot Run Test on env `%s`, allowed: %v", c.Environment.Stage, c.Environment.RunTestEnvironment))
	}
	return c, v, nil
}

func newHandler(t *testing.T) (*handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	c, v, err := newConfig(t)
	if err != nil {
		return nil, err
	}

	h.SetDefaultConfig(c)
	h.SetViper("config", v)

	if assert.NoError(t, initialize.LoadAllDatabaseConnection(h), "Error while loading all database connection: LoadAllDatabaseConnection") {
	}

	return h, nil
}
