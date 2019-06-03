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
var cmdtype int

const (
	// CmdLoginType 登录命令
	CmdLoginType = iota
	// CmdSftpType SFTP命令
	CmdSftpType
)

// Command 处理用户输入的命令
// Q/q 退出
// N/n 输出节点
// 其它命令则尝试匹配节点
// 		如果是直接回车，则输出所有节点列表
// 		如果匹配失败，则输出空节点列表
func Command() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("D>>>  ")
		cmd, _ := reader.ReadString('\n')
		cmd = textFormate(cmd)
		switch cmd {
		case "Q":
			fallthrough
		case "q":
			return
		// case "F":
		// 	fallthrough
		// case "f":
		// 	cmdtype = CmdSftpType
		// 	showNodes()
		case "N":
			fallthrough
		case "n":
			cmdtype = CmdLoginType
			showNodes()
		default:
			handlerUserCmd(cmd)
		}
	}
}

func textFormate(text string) string {
	text = strings.Replace(text, "\n", "", -1)
	return strings.TrimSpace(text)
}

func handlerUserCmd(key string) {
	outputMode()
	if idx, err := strconv.Atoi(key); err == nil {
		// 处理用户选中的节点ID
		if idx > len(chooseRNS)-1 {
			showNodes()
			return
		}

		switch cmdtype {
		case CmdLoginType:
			loginNode(chooseRNS[idx])
		case CmdSftpType:
			sftpNode(chooseRNS[idx])
		}

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

func sftpNode(rn model.RemoteNode) {
	err := node.Sftp(rn)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func loginNode(rn model.RemoteNode) {
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
		outputMode()
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

func outputMode() {
	switch cmdtype {
	case CmdLoginType:
		fmt.Println("Login Model!")
	case CmdSftpType:
		fmt.Println("SFTP Model!")
	}
}
