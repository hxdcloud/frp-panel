package controller

import (
	"frp-panel/config"
	"frp-panel/global"
	"frp-panel/model"
	"frp-panel/response"
	"frp-panel/utils"
	"github.com/fatedier/frp/pkg/metrics/mem"
	"github.com/fatedier/frp/pkg/util/version"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

type serverInfoResp struct {
	Version               string `json:"version"`
	BindPort              int    `json:"bind_port"`
	BindUDPPort           int    `json:"bind_udp_port"`
	VhostHTTPPort         int    `json:"vhost_http_port"`
	VhostHTTPSPort        int    `json:"vhost_https_port"`
	KCPBindPort           int    `json:"kcp_bind_port"`
	TcpmuxHttpConnectPort int    `json:"tcpmux_httpconnect_port"`
	SubdomainHost         string `json:"subdomain_host"`
	MaxPoolCount          int64  `json:"max_pool_count"`
	MaxPortsPerClient     int64  `json:"max_ports_per_client"`
	HeartBeatTimeout      int64  `json:"heart_beat_timeout"`

	TotalTrafficIn  int64            `json:"total_traffic_in"`
	TotalTrafficOut int64            `json:"total_traffic_out"`
	CurConns        int64            `json:"cur_conns"`
	ClientCounts    int64            `json:"client_counts"`
	ProxyTypeCounts map[string]int64 `json:"proxy_type_count"`

	Host      *model.Host      `json:"host"`
	HostState *model.HostState `json:"host_state"`
}

type ProxyStatsInfo struct {
	Name            string      `json:"name"`
	Conf            interface{} `json:"conf"`
	TodayTrafficIn  int64       `json:"today_traffic_in"`
	TodayTrafficOut int64       `json:"today_traffic_out"`
	CurConns        int64       `json:"cur_conns"`
	LastStartTime   string      `json:"last_start_time"`
	LastCloseTime   string      `json:"last_close_time"`
	Status          string      `json:"status"`
}

// GetFrpsConfig @Summary 获取frps配置信息
// @Schemes
// @Description get frps config
// @Tags frps
// @Accept json
// @Produce json
// @Success 200 {object} config.FrpsConfig
// @Router /api/frps/config [get]
func GetFrpsConfig(ctx *gin.Context) {
	response.Success(ctx, global.FRPS_CONFIG)
}

// GetFrpsInfo @Summary 获取frps信息
// @Schemes
// @Description get frps info
// @Tags frps
// @Accept json
// @Produce json
// @Success 200 {object} serverInfoResp
// @Router /api/frps/info [get]
func GetFrpsInfo(ctx *gin.Context) {
	serverStats := mem.StatsCollector.GetServer()
	svrResp := serverInfoResp{
		Version:               version.Full(),
		BindPort:              global.FRPS_CONFIG.Common.BindPort,
		BindUDPPort:           global.FRPS_CONFIG.Common.BindUdpPort,
		VhostHTTPPort:         global.FRPS_CONFIG.Common.VhostHttpPort,
		VhostHTTPSPort:        global.FRPS_CONFIG.Common.VhostHttpsPort,
		KCPBindPort:           global.FRPS_CONFIG.Common.KcpBindPort,
		TcpmuxHttpConnectPort: global.FRPS_CONFIG.Common.TcpmuxHttpConnectPort,
		SubdomainHost:         global.FRPS_SRV.Cfg.SubDomainHost,
		MaxPoolCount:          global.FRPS_SRV.Cfg.MaxPoolCount,
		MaxPortsPerClient:     global.FRPS_SRV.Cfg.MaxPortsPerClient,
		HeartBeatTimeout:      global.FRPS_SRV.Cfg.HeartbeatTimeout,

		TotalTrafficIn:  serverStats.TotalTrafficIn,
		TotalTrafficOut: serverStats.TotalTrafficOut,
		CurConns:        serverStats.CurConns,
		ClientCounts:    serverStats.ClientCounts,
		ProxyTypeCounts: serverStats.ProxyTypeCounts,

		Host:      utils.GetHost(),
		HostState: utils.GetState(),
	}
	response.Success(ctx, svrResp)
}

// SaveFrpsConfig @Summary 保存 frps 配置信息
// @Schemes
// @Description save frps config
// @Tags frps
// @Accept json
// @Produce json
// @Success 200
// @Router /api/frps/config [post]
func SaveFrpsConfig(ctx *gin.Context) {
	var cfg config.FrpsConfig
	_ = ctx.ShouldBindJSON(&cfg)
	global.FRPS_CONFIG = &cfg
	if global.FRPS_CONFIG.Common.BindAddr != "" {
		global.FRPS_CONFIG_VIPER.Set("common.bind_addr", global.FRPS_CONFIG.Common.BindAddr)
	}
	if global.FRPS_CONFIG.Common.BindPort != 0 {
		global.FRPS_CONFIG_VIPER.Set("common.bind_port", strconv.Itoa(global.FRPS_CONFIG.Common.BindPort))
	}
	err := global.FRPS_CONFIG_VIPER.WriteConfig()
	if err != nil {
		return
	}
	response.Success(ctx, nil)
	go os.Exit(3)
}
