package IPUtil

import (
	"errors"
	"net"
)

// 通过cidr 获取ip和网段地址
func GetIPNetByCIDR(ip string, cidr int) (net.IP, *net.IPNet, error) {
	ipv4 := net.ParseIP(ip).To4()
	mask := net.CIDRMask(cidr, 32)

	ipnet := &net.IPNet{IP: ipv4, Mask: mask}

	return net.ParseCIDR(ipnet.String())
}

// 通过mask地址获取网段
func GetIPNetByMask(ip string, mask string) (net.IP, *net.IPNet, error) {
	ipv4 := net.ParseIP(ip).To4()
	mkv34 := net.ParseIP(mask).To4()

	if len(ipv4) != net.IPv4len || len(mkv34) != int(net.IPv4len) {
		return nil, nil, errors.New("不正确的ip或掩码地址")
	}

	ipnet := &net.IPNet{
		IP:   ipv4,
		Mask: net.IPv4Mask(mkv34[0], mkv34[1], mkv34[2], mkv34[3]),
	}

	return net.ParseCIDR(ipnet.String())

}

// 获取所有IP信息
func GetIpsByCIDR(cidr string) []string {
	// 解析CIDR表达式
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil
	}

	inc := func(p net.IP) {
		for i := len(p) - 1; i >= 0; i-- {
			p[i]++
			if p[i] > 0 {
				break
			}
		}
	}
	// 计算网络内的所有IP地址
	var ips []string
	for ip := network.IP.Mask(network.Mask); network.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips
}
