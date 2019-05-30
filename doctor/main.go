package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/andy-zhangtao/doctor/doctor/cli"
	"github.com/andy-zhangtao/doctor/doctor/rpc/impl"
	"github.com/andy-zhangtao/doctor/doctor/server"
)

/**
doctor 是堡垒机的核心组件

|--------|      |-------|       |------|
|  USER  |----->| Doctor|------>| Node |
|--------|      |-------|       |------|

doctor会尝试将用户的登录请求转发到远程节点。
**/

func main() {
	isServer := flag.Bool("server", false, "Run as daemon mode")
	isCli := flag.Bool("client", false, "Run as client mode")
	logLevel := flag.String("level", "debug", "Log Level, debug/warn/error/panic")
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

	if *isServer {
		fmt.Println("===Server Mode===")
		go func() {
			if err := impl.Run(); err != nil {
				fmt.Printf("GRPC Error: %s \n", err)
			}
		}()
		server.StartWeb()
	}

	if *isCli {
		cli.Command()
	}
}
