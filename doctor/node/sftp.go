package node

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"github.com/pkg/sftp"

	"github.com/andy-zhangtao/doctor/doctor/model"
	"golang.org/x/crypto/ssh"
)

const destDir = "/tmp"
const nurseBIN = "nurse"

const (
	directUp = iota
	directDownload
)

func textFormate(text string) string {
	text = strings.Replace(text, "\n", "", -1)
	return strings.TrimSpace(text)
}

// Sftp 处理远程节点的SFTP请求
func Sftp(rn model.RemoteNode) (err error) {
	if err = sftpLogin(rn); err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("SFTP>  ")
		cmd, _ := reader.ReadString('\n')
		cmd = textFormate(cmd)
		switch cmd {
		case "bye":
			fallthrough
		case "quit":
			fallthrough
		case "exit":
			return
		default:
			executeSFTPCmd(rn, cmd)
		}
	}
	return
}

func sftpLogin(rn model.RemoteNode) (err error) {
	signer, err := ssh.ParsePrivateKey([]byte(rn.Key))
	if err != nil {
		return fmt.Errorf("unable to parse private key: %v", err)
	}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", rn.Ip), &ssh.ClientConfig{
		User: rn.Name,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			// fmt.Printf("Login %s As %s Use Private Key \n", rn.Ip, rn.Name)
			return nil
		},
	})

	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	pwd, err := client.Getwd()
	if err != nil {
		return
	}

	fmt.Println(pwd)
	return
}

func executeSFTPCmd(rn model.RemoteNode, cmd string) {
	if strings.HasPrefix(cmd, "!") {
		executeLocalCmd(cmd)
		return
	}

	executeRemoteCmd(rn, cmd)
	return
}

func executeLocalCmd(cmd string) {

}

func executeRemoteCmd(rn model.RemoteNode, cmd string) {
	if err := command(rn, []string{cmd}); err != nil {
		fmt.Println(err)
	}
}

// transFileToNode 向远程节点推送指定文件
// 默认推送到远程节点的/tmp目录中
func transFileToNode(rn model.RemoteNode, files []string, direct int) (err error) {

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", rn.Ip), &ssh.ClientConfig{
		User: rn.Name,
		Auth: []ssh.AuthMethod{ssh.Password(rn.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			fmt.Printf("Login %s As %s \n", rn.Ip, rn.Name)
			return nil
		},
	})

	if err != nil {
		return
	}

	remoteSFTP, err := sftp.NewClient(client)
	if err != nil {
		return
	}

	defer remoteSFTP.Close()

	switch direct {
	case directUp:
		for _, name := range files {
			destFilePath := fmt.Sprintf("%s/d_%s", destDir, name)
			originFilePath := fmt.Sprintf("%s/%s", destDir, name)
			bin, err := remoteSFTP.Create(destFilePath)
			if err != nil {
				return err
			}

			origin, err := ioutil.ReadFile(originFilePath)
			if err != nil {
				return err
			}

			originFileInfo, err := os.Stat(originFilePath)
			if err != nil {
				return err
			}

			if _, err = bin.Write(origin); err != nil {
				return err
			}

			f, err := remoteSFTP.Lstat(destFilePath)
			if err != nil {
				return err
			}

			if f.Size() != int64(len(origin)) {
				return errors.New("Nurse Size Not Compare! ")
			}

			err = bin.Chmod(originFileInfo.Mode().Perm())
			if err != nil {
				return err
			}
		}
	case directDownload:
		// for _, name := range files {
		// 	filePath := fmt.Sprintf("%s/%s", destDir, name)
		// 	bin, err := remoteSFTP.Create(filePath)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	origin, err := ioutil.ReadFile(filePath)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	if _, err = bin.Write(origin); err != nil {
		// 		return err
		// 	}

		// 	f, err := remoteSFTP.Lstat(filePath)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	if f.Size() != int64(len(origin)) {
		// 		return errors.New("Nurse Size Not Compare! ")
		// 	}
		// }
	}

	return
}
