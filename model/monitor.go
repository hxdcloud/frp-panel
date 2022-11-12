package model

type Host struct {
	Platform        string   `json:"platform"`
	PlatformVersion string   `json:"platform_version"`
	CPU             []string `json:"cpu"`
	MemTotal        uint64   `json:"mem_total"`
	SwapTotal       uint64   `json:"swap_total"`
	Arch            string   `json:"arch"`
	Virtualization  string   `json:"virtualization"`
	BootTime        uint64   `json:"boot_time"`
}
type HostState struct {
	CPU            float64 `json:"cpu"`
	MemPercent     float64 `json:"mem_percent"`
	MemUsed        uint64  `json:"mem_used"`
	SwapPercent    float64 `json:"swap_percent"`
	SwapUsed       uint64  `json:"swap_used"`
	NetInTransfer  uint64  `json:"net_in_transfer"`
	NetOutTransfer uint64  `json:"net_out_transfer"`
	NetInSpeed     uint64  `json:"net_in_speed"`
	NetOutSpeed    uint64  `json:"net_out_speed"`
	Uptime         uint64  `json:"uptime"`
	Load1          float64 `json:"load_1"`
	Load5          float64 `json:"load_5"`
	Load15         float64 `json:"load_15"`
	TcpConnCount   uint64  `json:"tcp_conn_count"`
	UdpConnCount   uint64  `json:"udp_conn_count"`
	ProcessCount   uint64  `json:"process_count"`
}
