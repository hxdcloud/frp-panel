package model

import (
	"github.com/fatedier/frp/pkg/msg"
)

type ProxyReq struct {
	Proxy msg.NewProxy `json:"proxy"`
	RunID string       `json:"runID"`
}
