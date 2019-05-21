package main

import (
	"flag"
	"fmt"

	"github.com/andy-zhangtao/doctor/nurse/rpc"
)

// nurse 远程节点初始化
// 1. 安装登录用户
// 2. 安装ssh key
// 3. 采集节点信息
// 4. 通知Doctor节点初始化结果
func main() {

	serverAddress := flag.String("server", "", "Doctor Server Address")

	flag.Parse()

	if *serverAddress == "" {
		flag.Usage()
		return
	}

	if err := rpc.ReportNodeInfo(*serverAddress); err != nil {
		fmt.Println(err)
	}
}
