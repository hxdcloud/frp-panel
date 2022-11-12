package legocmd_test

import (
	"frp-panel/common/legocmd"
	"testing"
)

func TestLegoClient(t *testing.T) {
	_, err := legocmd.New()
	if err != nil {
		t.Error(err)
	}
}

func TestLegoCloudflareDNSCert(t *testing.T) {
	lego, err := legocmd.New()
	if err != nil {
		t.Error(err)
	}
	var (
		domain   string = "lego.365188.xyz"
		email    string = "wangxingrui1997@gmail.com"
		provider string = "cloudflare"
		DNSEnv   map[string]string
	)
	DNSEnv = make(map[string]string)
	//DNSEnv["CLOUDFLARE_EMAIL"] = "wangxingrui1997@gmail.com"
	DNSEnv["CLOUDFLARE_DNS_API_TOKEN"] = "mL3vhXWUHtbHlolcbQBrTYyFMT0kokQOzgs-aBQD"
	certPath, keyPath, err := lego.DNSCert(domain, email, provider, DNSEnv)
	if err != nil {
		t.Error(err)
	}
	t.Log(certPath)
	t.Log(keyPath)
}

func TestLegoDNSCert(t *testing.T) {
	lego, err := legocmd.New()
	if err != nil {
		t.Error(err)
	}
	var (
		domain   string = "node1.test.com"
		email    string = "test@gmail.com"
		provider string = "alidns"
		DNSEnv   map[string]string
	)
	DNSEnv = make(map[string]string)
	DNSEnv["ALICLOUD_ACCESS_KEY"] = "aaa"
	DNSEnv["ALICLOUD_SECRET_KEY"] = "bbb"
	certPath, keyPath, err := lego.DNSCert(domain, email, provider, DNSEnv)
	if err != nil {
		t.Error(err)
	}
	t.Log(certPath)
	t.Log(keyPath)
}

func TestLegoHTTPCert(t *testing.T) {
	lego, err := legocmd.New()
	if err != nil {
		t.Error(err)
	}
	var (
		domain string = "node1.test.com"
		email  string = "test@gmail.com"
	)
	certPath, keyPath, err := lego.HTTPCert(domain, email)
	if err != nil {
		t.Error(err)
	}
	t.Log(certPath)
	t.Log(keyPath)
}

func TestLegoRenewCert(t *testing.T) {
	lego, err := legocmd.New()
	if err != nil {
		t.Error(err)
	}
	var (
		domain   string = "node1.test.com"
		email    string = "test@gmail.com"
		provider string = "alidns"
		DNSEnv   map[string]string
	)
	DNSEnv = make(map[string]string)
	DNSEnv["ALICLOUD_ACCESS_KEY"] = "aaa"
	DNSEnv["ALICLOUD_SECRET_KEY"] = "bbb"
	certPath, keyPath, err := lego.RenewCert(domain, email, "dns", provider, DNSEnv)
	if err != nil {
		t.Error(err)
	}
	t.Log(certPath)
	t.Log(keyPath)

	certPath, keyPath, err = lego.RenewCert(domain, email, "http", provider, DNSEnv)
	if err != nil {
		t.Error(err)
	}
	t.Log(certPath)
	t.Log(keyPath)
}
