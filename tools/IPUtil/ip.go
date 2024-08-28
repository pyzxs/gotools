package IPUtil

import (
	"fmt"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

func Ip2Region(ip string) string {

	searcher, err := xdb.NewWithFileOnly(GetXdbPath() + "ip2region.xdb")
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return ""
	}

	defer searcher.Close()

	str, err := searcher.SearchByStr(ip)
	if err != nil {
		return ""
	}
	arr := strings.Split(str, "|")

	if len(arr) < 3 {
		return arr[0]
	}
	return fmt.Sprintf("%s-%s-%s", arr[0], arr[2], arr[3])
}
