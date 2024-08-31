package IPUtil

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

type Ip2Region struct {
	XDbFilePath string // xdb地址
}

func (i *Ip2Region) GetRegion(ip string) (string, error) {

	searcher, err := xdb.NewWithFileOnly(i.XDbFilePath)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return "", err
	}

	defer searcher.Close()

	str, err := searcher.SearchByStr(ip)
	if err != nil {
		return "", err
	}
	arr := strings.Split(str, "|")

	if len(arr) < 3 {
		return arr[0], nil
	}

	return fmt.Sprintf("%s-%s-%s", arr[0], arr[2], arr[3]), nil
}

// ip地址转无符号数值
func Ipv4ToLong(ipStr string) uint32 {
	ip := net.ParseIP(ipStr).To4()
	if ip == nil {
		return 0
	}

	return binary.BigEndian.Uint32(ip)
}

// 无符号32位数值转IP地址
func LongToIpv4(long uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, long)
	ip := net.IP(ipByte)
	return ip.String()
}
