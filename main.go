package main

import (
	"frp-panel/common/frps"
	"frp-panel/core"
	"frp-panel/global"
	"frp-panel/routers"
)

// @title frp-panel
// @version 1.0 版本
// @description frp-panel API 描述
// @BasePath /api  基础路径
// @query.collection.format multi
func main() {

	global.FRPS_CONFIG_VIPER = core.FrpsConfigViper() // 初始化FrpsConfigViper
	global.FRPC_CONFIG_VIPER = core.FrpcConfigViper() // 初始化FrpcConfigViper
	global.VALIDATE = core.Validator()                // 初始化Validate
	go frps.Execute()                                 // 启动frps
	r := routers.SetupRouter()
	r.Run()
}
