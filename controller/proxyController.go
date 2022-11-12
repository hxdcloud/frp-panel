package controller

import (
	"frp-panel/global"
	"frp-panel/response"
	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/server"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

var (
	TCP = "tcp"
)

// GetTcpProxy @Summary 获取frps TCP 代理信息
// @Schemes
// @Description get frps tcp proxies
// @Tags proxy
// @Accept json
// @Produce json
// @Param current query int true "current page"
// @Param pageSize query int true "page size"
// @Param name query int false "proxy name"
// @Param status query string false "proxy status"
// @Success 200
// @Router /api/proxy/tcp [get]
func GetTcpProxy(ctx *gin.Context) {
	name := ctx.Query("name")
	status := ctx.Query("status")

	proxyInfoResp := server.GetProxyInfoResp{}
	proxyInfoResp.Proxies = global.FRPS_SRV.GetProxyStatsByType("tcp")

	if name != "" {
		for i := 0; i < len(proxyInfoResp.Proxies); {
			if !strings.Contains(proxyInfoResp.Proxies[i].Name, name) {
				proxyInfoResp.Proxies = append(proxyInfoResp.Proxies[:i], proxyInfoResp.Proxies[i+1:]...)
			} else {
				i++
			}
		}
	}

	if status != "" {
		for i := 0; i < len(proxyInfoResp.Proxies); {
			if proxyInfoResp.Proxies[i].Status != status {
				proxyInfoResp.Proxies = append(proxyInfoResp.Proxies[:i], proxyInfoResp.Proxies[i+1:]...)
			} else {
				i++
			}
		}
	}

	response.PageDataSuccess(ctx, proxyInfoResp.Proxies, len(proxyInfoResp.Proxies))
}

// GetUdpProxy @Summary 获取frps UDP 代理信息
// @Schemes
// @Description get frps udp proxies
// @Tags proxy
// @Accept json
// @Produce json
// @Success 200
// @Router /api/proxy/udp [get]
func GetUdpProxy(ctx *gin.Context) {
	proxyInfoResp := server.GetProxyInfoResp{}
	proxyInfoResp.Proxies = global.FRPS_SRV.GetProxyStatsByType("udp")
	response.Success(ctx, proxyInfoResp.Proxies)
}

// GetHttpProxy @Summary 获取frps HTTP 代理信息
// @Schemes
// @Description get frps http proxies
// @Tags proxy
// @Accept json
// @Produce json
// @Success 200
// @Router /api/proxy/http [get]
func GetHttpProxy(ctx *gin.Context) {
	proxyInfoResp := server.GetProxyInfoResp{}
	proxyInfoResp.Proxies = global.FRPS_SRV.GetProxyStatsByType("http")
	response.Success(ctx, proxyInfoResp.Proxies)
}

// GetHttpsProxy @Summary 获取frps HTTPS 代理信息
// @Schemes
// @Description get frps https proxies
// @Tags proxy
// @Accept json
// @Produce json
// @Success 200
// @Router /api/proxy/https [get]
func GetHttpsProxy(ctx *gin.Context) {
	proxyInfoResp := server.GetProxyInfoResp{}
	proxyInfoResp.Proxies = global.FRPS_SRV.GetProxyStatsByType("https")
	response.Success(ctx, proxyInfoResp.Proxies)
}

// GetStcpProxy @Summary 获取frps STCP 代理信息
// @Schemes
// @Description get frps stcp proxies
// @Tags proxy
// @Accept json
// @Produce json
// @Success 200
// @Router /api/proxy/stcp [get]
func GetStcpProxy(ctx *gin.Context) {
	proxyInfoResp := server.GetProxyInfoResp{}
	proxyInfoResp.Proxies = global.FRPS_SRV.GetProxyStatsByType("stcp")
	response.Success(ctx, proxyInfoResp.Proxies)
}

// GetSudpProxy @Summary 获取frps SUDP 代理信息
// @Schemes
// @Description get frps sudp proxies
// @Tags proxy
// @Accept json
// @Produce json
// @Success 200
// @Router /api/proxy/sudp [get]
func GetSudpProxy(ctx *gin.Context) {
	proxyInfoResp := server.GetProxyInfoResp{}
	proxyInfoResp.Proxies = global.FRPS_SRV.GetProxyStatsByType("sudp")
	response.Success(ctx, proxyInfoResp.Proxies)
}

// GetXtcpProxy @Summary 获取frps XTCP 代理信息
// @Schemes
// @Description get frps xtcp proxies
// @Tags proxy
// @Accept json
// @Produce json
// @Success 200
// @Router /api/proxy/xtcp [get]
func GetXtcpProxy(ctx *gin.Context) {
	proxyInfoResp := server.GetProxyInfoResp{}
	proxyInfoResp.Proxies = global.FRPS_SRV.GetProxyStatsByType("xtcp")
	response.Success(ctx, proxyInfoResp.Proxies)
}

// GetTcpmuxProxy @Summary 获取frps TCPMUX 代理信息
// @Schemes
// @Description get frps tcpmux proxies
// @Tags proxy
// @Accept json
// @Produce json
// @Success 200
// @Router /api/proxy/tcpmux [get]
func GetTcpmuxProxy(ctx *gin.Context) {
	proxyInfoResp := server.GetProxyInfoResp{}
	proxyInfoResp.Proxies = global.FRPS_SRV.GetProxyStatsByType("tcpmux")
	response.Success(ctx, proxyInfoResp.Proxies)
}

// AddTcpProxy @Summary 新增TCP代理
// @Schemes
// @Description add frp tcp proxy
// @Tags proxy
// @Accept json
// @Produce json
// @Param pxy body msg.NewProxyIni true "proxy info"
// @Success 200
// @Router /api/proxy/tcp [post]
func AddTcpProxy(ctx *gin.Context) {

	// TODO 校验参数

	var pxy = msg.NewProxyIni{}
	_ = ctx.ShouldBindJSON(&pxy)

	ctl, exist := global.FRPS_SRV.CtlManager.GetByID(pxy.RunId)
	if !exist {
		response.Fail(ctx, nil)
		return
	}

	pxy.ProxyType = TCP
	ctl.SendCh <- &pxy

	response.Success(ctx, pxy)
}

// DeleteProxy @Summary 删除TCP代理
// @Schemes
// @Description delete frp proxy
// @Tags proxy
// @Accept json
// @Produce json
// @Param pxy body msg.DeleteProxy true "proxy info"
// @Success 200
// @Router /api/proxy/tcp [delete]
func DeleteProxy(ctx *gin.Context) {

	var pxy = msg.DeleteProxy{}
	_ = ctx.ShouldBindJSON(&pxy)

	ctl, exist := global.FRPS_SRV.CtlManager.GetByID(pxy.RunId)
	if !exist {
		response.Fail(ctx, nil)
		return
	}

	ctl.SendCh <- &pxy

	response.Success(ctx, pxy)
}

// ValidateProxyName @Summary 校验代理名称是否可用
// @Schemes
// @Description validate frp proxy name unique
// @Tags proxy
// @Accept json
// @Produce json
// @Param name query string true "proxy name"
// @Param runId query string true "frpc runId"
// @Success 200
// @Router /api/proxy/name/validate [get]
func ValidateProxyName(ctx *gin.Context) {
	name := ctx.Query("name")
	runId := ctx.Query("runId")

	ctl, exist := global.FRPS_SRV.CtlManager.GetByID(runId)
	if !exist {
		response.Success(ctx, false)
		return
	}
	for _, proxy := range ctl.PxyManager.Pxys {
		if proxy.GetName() == name {
			response.Success(ctx, false)
		}
	}
	response.Success(ctx, true)
}

// ValidateTcpPort @Summary 校验tcp端口是否可用
// @Schemes
// @Description validate frp tcp port usable
// @Tags proxy
// @Accept json
// @Produce json
// @Param port query int true "tcp port"
// @Success 200
// @Router /api/proxy/port/tcp/validate [get]
func ValidateTcpPort(ctx *gin.Context) {
	port, err := strconv.Atoi(ctx.Query("port"))
	if err != nil {
		return
	}
	response.Success(ctx, global.FRPS_SRV.Rc.TCPPortManager.IsPortAvailable(port))
}
