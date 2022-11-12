package core

import (
	"fmt"
	"frp-panel/core/internal"
	"frp-panel/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func FrpsConfigViper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("/")
	v.SetConfigFile(internal.FrpsConfigFile)
	v.SetConfigType("ini")

	v.SetDefault("common.bind_port", 7000)

	err := v.ReadInConfig()
	if err != nil {
		os.Create(internal.FrpsConfigFile)
		v.WriteConfig()
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.FRPS_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.FRPS_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}

func FrpcConfigViper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("/")
	v.SetConfigFile(internal.FrpcConfigFile)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		os.Create(internal.FrpcConfigFile)
		v.WriteConfig()
	}

	v.WatchConfig()
	//v.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("config file changed:", e.Name)
	//	if err = v.Unmarshal(&global.FRPS_CONFIG); err != nil {
	//		fmt.Println(err)
	//	}
	//})
	//if err = v.Unmarshal(&global.FRPS_CONFIG); err != nil {
	//	fmt.Println(err)
	//}
	return v
}
