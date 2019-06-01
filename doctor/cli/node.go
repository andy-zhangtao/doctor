package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/andy-zhangtao/doctor/doctor/model"
	"github.com/andy-zhangtao/doctor/doctor/node"
)

var err error
var rns []model.RemoteNode
var chooseRNS []model.RemoteNode

// Command 处理用户输入的命令
// Q/q 退出
// N/n 输出节点
// 其它命令则尝试匹配节点
// 		如果是直接回车，则输出所有节点列表
// 		如果匹配失败，则输出空节点列表
func Command() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">>>  ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		switch cmd {
		case "Q":
			fallthrough
		case "q":
			return
		case "N":
			fallthrough
		case "n":
			showNodes()
		default:
			handlerUserCmd(cmd)
		}
	}
}

func handlerUserCmd(key string) {
	// if key == "" {
	// 	// 如果输入的是回车，则直接输出所有节点列表
	// 	showNodes()
	// 	return
	// }

	if idx, err := strconv.Atoi(key); err == nil {
		if idx > len(chooseRNS)-1 {
			showNodes()
			return
		}
		loginNode(chooseRNS[idx])

	}
	if len(rns) == 0 {
		rns, err = getAllNodes()
		if err != nil {
			fmt.Printf("%#v \n", err)
			return
		}
	}

	var _chooseRNS []model.RemoteNode
	for _, r := range rns {
		if strings.Contains(r.Ip, key) || strings.Contains(r.Comment, key) {
			_chooseRNS = append(_chooseRNS, r)
		}
	}

	chooseRNS = _chooseRNS
	showNodes()
	return
}

func loginNode(rn model.RemoteNode) {
	fmt.Printf("Login %s \n", rn.Ip)
	err := node.Login(rn)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func showNodes() {
	fmt.Println("---------------------------------------------------------------")
	fmt.Println()

	if len(rns) == 0 {
		rns, err = getAllNodes()
		if err != nil {
			fmt.Printf("%#v \n", err)
			return
		}
	}

	if len(chooseRNS) == 0 {
		fmt.Println("No Match Node!")
	} else {
		for i, r := range chooseRNS {
			fmt.Printf("[%2d] | %15s | %10s | %20s \n", i, r.Ip, r.Name, r.Comment)
		}
	}

	fmt.Println()
}

func getAllNodes() (rns []model.RemoteNode, err error) {
	return node.FindAllNodes()
}

func showHead() {
	fmt.Println()
	fmt.Println("========DOCTOR========")
	fmt.Println("")
	fmt.Println("   1> 输入 N/n 显示所有节点 ")
	fmt.Println("   2> 输入 F/f 上传/下载文件 ")
	fmt.Println("   3> 输入 Q/q 退出Doctor ")
	fmt.Println("")

}
