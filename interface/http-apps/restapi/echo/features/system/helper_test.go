package system

import (
	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/spf13/viper"
)

func newConfig(h *handler.Handler) (*config.Config, *viper.Viper, error) {
	configPath := "../../../../../../conf"

	//init config
	cfg, viper, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}
	cfg.IAM.Casbin.ModelPath = "../../../../../../conf/casbin/casbin_rbac_rest_model.conf"

	h.SetDefaultConfig(cfg)

	return cfg, viper, nil
}

func newHandler() *handler.Handler {

	h, _ := handler.NewHandler()

	// init configuration
	_, _, err := newConfig(h)
	if err != nil {
		panic(err)
	}

	return h
}
