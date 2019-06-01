package main

import (
	"flag"
	"strings"

	"github.com/andy-zhangtao/doctor/nurse/node"
	"github.com/andy-zhangtao/doctor/nurse/rpc"
	"github.com/sirupsen/logrus"
)

// nurse 远程节点初始化
// 1. 安装登录用户
// 2. 安装ssh key
// 3. 采集节点信息
// 4. 通知Doctor节点初始化结果
func main() {

	serverAddress := flag.String("server", "", "Doctor Server Address")
	myIP := flag.String("ip", "", "My IP. Use for register node ")
	logLevel := flag.String("level", "debug", "Nurse Log Level debug/warn/error/painc")

	flag.Parse()

	switch strings.ToLower(*logLevel) {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "panic":
		logrus.SetLevel(logrus.ErrorLevel)
	}

	if *serverAddress == "" {
		flag.Usage()
		return
	}

	var key []byte
	var err error
	if key, err = node.InitRemoteNode(); err != nil {
		logrus.Errorf("Node Init Error: [%s]", err.Error())
		return
	}

	if err := rpc.ReportNodeInfo(*serverAddress, *myIP, key); err != nil {
		logrus.Errorf("Node Register Error: [%s]", err.Error())
	}
	return
}
