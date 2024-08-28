//go:build !windows
// +build !windows

package IPUtil

import (
	"fmt"
	"os"
)

func GetXdbPath() string {
	gopath := os.Getenv("GOPATH")
	project := "github.com/pyzxs/gotools/tools/IPUtil/"
	return fmt.Sprintf("%s%s%s", gopath, "/pkg/mod", project)
}
