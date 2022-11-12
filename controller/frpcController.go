package controller

import (
	"frp-panel/global"
	"frp-panel/model"
	"frp-panel/response"
	"github.com/fatedier/frp/pkg/msg"
	"github.com/gin-gonic/gin"
	"time"
)

type clientInfoResp struct {
	Status       string     `json:"status"`
	LoginMsg     *msg.Login `json:"loginMsg"`
	ProxiesCount int        `json:"proxiesCount"`
	LastPing     time.Time  `json:"lastPing"`
}

// GetFrpc @Summary 获取frpc
// @Schemes
// @Description get frpc
// @Tags frpc
// @Accept json
// @Produce json
// @Success 200
// @Router /api/frpc [get]
func GetFrpc(ctx *gin.Context) {
	controllers := global.FRPS_SRV.CtlManager.CtlsByRunID
	frpcInfos := make([]*clientInfoResp, 0, len(controllers))
	for _, rid := range controllers {
		frpcInfo := &clientInfoResp{}
		controller, _ := global.FRPS_SRV.CtlManager.GetByID(rid.RunID)
		frpcInfo.LoginMsg = controller.LoginMsg
		frpcInfo.Status = controller.Status
		frpcInfo.ProxiesCount = len(controller.PxyManager.Pxys)
		frpcInfo.LastPing = rid.LastPing
		frpcInfos = append(frpcInfos, frpcInfo)
	}
	response.Success(ctx, frpcInfos)
}

// GetFrpcOption @Summary 获取frpc option
// @Schemes
// @Description get frpc option
// @Tags frpc
// @Accept json
// @Produce json
// @Success 200
// @Router /api/frpc/option [get]
func GetFrpcOption(ctx *gin.Context) {
	controllers := global.FRPS_SRV.CtlManager.CtlsByRunID
	options := make([]*model.Option, 0, len(controllers))
	for _, rid := range controllers {
		option := &model.Option{}
		controller, _ := global.FRPS_SRV.CtlManager.GetByID(rid.RunID)
		option.Label = controller.LoginMsg.Metas["name"]
		option.Value = controller.LoginMsg.RunID
		options = append(options, option)
	}
	response.Data(ctx, options)
}
