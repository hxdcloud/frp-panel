package global

import (
	"frp-panel/config"
	"github.com/fatedier/frp/server"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var (
	FRPS_CONFIG_VIPER *viper.Viper
	FRPC_CONFIG_VIPER *viper.Viper
	FRPS_CONFIG       *config.FrpsConfig
	FRPS_SRV          *server.Service
	VALIDATE          *validator.Validate
)
