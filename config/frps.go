package config

type FrpsConfig struct {
	Common Common `mapstructure:"common" json:"common" structs:"common"`
}

type Common struct {
	// bind config
	BindAddr       string `mapstructure:"bind_addr" json:"bindAddr,omitempty"`
	ProxyBindAddr  string `mapstructure:"proxy_bind_addr" json:"proxyBindAddr"`
	BindPort       int    `mapstructure:"bind_port" json:"bindPort,omitempty" validate:"required"`
	BindUdpPort    int    `mapstructure:"bind_udp_port" json:"bindUdpPort,omitempty"`
	KcpBindPort    int    `mapstructure:"kcp_bind_port" json:"KcpBindPort,omitempty"`
	VhostHttpPort  int    `mapstructure:"vhost_http_port" json:"vhostHttpPort,omitempty"`
	VhostHttpsPort int    `mapstructure:"vhost_https_port" json:"vhostHttpsPort,omitempty"`

	// connect config
	TcpKeepalive     int `mapstructure:"tcp_keepalive" json:"tcpKeepalive"`
	HeartbeatTimeout int `mapstructure:"heartbeat_timeout" json:"heartbeatTimeout"`
	UserConnTimeout  int `mapstructure:"user_conn_timeout" json:"userConnTimeout"`
	UdpPacketSize    int `mapstructure:"udp_packet_size" json:"udpPacketSize"`

	// auth config
	AuthenticationMethod     string `mapstructure:"authentication_method" json:"authenticationMethod"`
	AuthenticateHeartbeats   bool   `mapstructure:"authenticate_heartbeats" json:"authenticateHeartbeats"`
	AuthenticateNewWorkConns bool   `mapstructure:"authenticate_new_work_conns" json:"authenticateNewWorkConns"`
	Token                    string `mapstructure:"token" json:"token"`
	OidcIssuer               string `mapstructure:"oidc_issuer" json:"OidcIssuer"`
	OidcAudience             string `mapstructure:"oidc_audience" json:"oidcAudience"`
	OidcSkipExpiryCheck      bool   `mapstructure:"oidc_skip_expiry_check" json:"oidcSkipExpiryCheck"`
	OidcSkipIssuerCheck      bool   `mapstructure:"oidc_skip_issuer_check" json:"oidcSkipIssuerCheck"`

	// manage config
	AllowPorts        string `mapstructure:"allow_ports" json:"allowPorts"`
	MaxPoolCount      int    `mapstructure:"max_pool_count" json:"maxPoolCount"`
	MaxPortsPerClient int    `mapstructure:"max_ports_per_client" json:"maxPortsPerClient"`
	TlsOnly           bool   `mapstructure:"tls_only" json:"tlsOnly"`

	// http & https config
	VhostHttpTimeout int    `mapstructure:"vhost_http_timeout" json:"vhostHttpTimeout"`
	SubdomainHost    string `mapstructure:"subdomain_host" json:"subdomainHost"`
	Custom404Page    string `mapstructure:"custom_404_page" json:"custom404Page"`

	// tcpmux config
	TcpmuxHttpConnectPort int  `mapstructure:"tcpmux_httpconnect_port" json:"tcpmuxHttpConnectPort"`
	TcpmuxPassthrough     bool `mapstructure:"tcpmux_passthrough" json:"tcpmuxPassthrough"`

	// log config
	LogFile                string `mapstructure:"log_file" json:"logFile,omitempty"`
	LogLevel               string `mapstructure:"log_level" json:"logLevel,omitempty"`
	LogMaxDays             int    `mapstructure:"log_max_days" json:"logMaxDays"`
	DisableLogColor        bool   `mapstructure:"disable_log_color" json:"disableLogColor"`
	DetailedErrorsToClient bool   `mapstructure:"detailed_errors_to_client" json:"detailedErrorsToClient"`
}
