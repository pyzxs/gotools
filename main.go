package main

import (
	"fmt"

	"github.com/pyzxs/gotools/tools/IPUtil"
)

func main() {
	region := IPUtil.Ip2Region{XDbFilePath: "./ip2region.xdb"}
	ips, err := region.GetRegion("153.3.238.110")
	if err != nil {
		panic(err)
	}

	fmt.Print(ips)
}
