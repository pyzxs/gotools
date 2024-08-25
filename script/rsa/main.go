package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	output string
	pri    string
	pub    string
)

func init() {
	flag.StringVar(&output, "o", "./", "输出地址")
	flag.StringVar(&pri, "pv", "rsa_private_key.pem", "私钥文件")
	flag.StringVar(&pub, "pu", "rsa_public_key", "公钥文件")
}

func createPrivateFile() {
	shell := fmt.Sprintf("openssl genrsa -out %s 1024", output+pri)
	cmd := Execute(shell)
	data, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}

	err = os.WriteFile(output+pri, data, 755)
	if err != nil {
		panic(err.Error())
	}
}

func createPublicFile() {
	shell := fmt.Sprintf("openssl rsa -in %s -pubout -out %s", output+pri, output+pub)
	cmd := Execute(shell)
	data, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}

	err = os.WriteFile(output+pri, data, 755)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	flag.Parse()
	if pri == "" {
		pri = "rsa_private_key.pem"
	}
	if pub == "" {
		pri = "rsa_public_key.pem"
	}
	fmt.Println("start create private file")
	createPrivateFile()
	fmt.Println("start create public file")
	createPublicFile()
}
