package IPUtil

import "testing"

// 通过CIDR获取网段信息
func TestGetIPNetByCIDR(t *testing.T) {
	ip, net, err := GetIPNetByCIDR("192.168.31.4", 24)
	t.Log(ip.String(), net.String(), err)
}

// 通过mask地址获取网段
func TestGetIPNetByMask(t *testing.T) {
	ip, net, err := GetIPNetByMask("192.168.31.4", "255.255.255.0")
	t.Log(ip.String(), net.String(), err)
}


// 获取所有IP信息
func TestGetIpsByCIDR(t *testing.T) {
	ips := GetIpsByCIDR("192.168.31.4/16")
	for _, p := range ips {
		t.Log(p)
	}
}
