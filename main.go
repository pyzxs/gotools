package main

import (
	"fmt"

	"github.com/pyzxs/gotools/tools/IPUtil"
)

func main() {
	region := IPUtil.Ip2Region("39.105.152.173")
	fmt.Println(region)
}
