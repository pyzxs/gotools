# IPUtil 项目包
- [IP2REGION](#IP2REGION)
- [IP地址转化](#IP地址转化)
- [IP网段](#IP网段)



## Usage

### IP2REGiON

下载ip2region.xdb文件

```go
region := IPUtil.Ip2Region{XDbFilePath: "./ip2region.xdb"}
ips, err := region.GetRegion("153.3.238.110")
if err != nil {
    panic(err)
}

fmt.Print(ips)

# 输出结果 中国-江苏省-南京市
```

### IP地址转化

```go
// IPv4地址转数值
func Ipv4ToLong(ip string) uin32
// 数值转IP地址
func LongToIpv4(long uint32) string 
```

### IP网段

```go
// 通过cidr 获取ip和网段地址
func GetIPNetByCIDR(ip string, cidr int) (net.IP, *net.IPNet, error)
// 通过mask地址获取网段
func GetIPNetByMask(ip string, mask string) (net.IP, *net.IPNet, error) 
// 获取网段内的所有IP地址
func GetIpsByCIDR(cidr string) []string 
```