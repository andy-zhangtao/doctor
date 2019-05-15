package node

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/pkg/sftp"

	"github.com/andy-zhangtao/doctor/model"
	"golang.org/x/crypto/ssh"
)

const destDir = "/tmp"
const nurseBIN = "nurse"

const (
	directUp = iota
	directDownload
)

// transFileToNode 向远程节点推送指定文件
// 默认推送到远程节点的/tmp目录中
func transFileToNode(rn model.RemoteNode, files []string, direct int) (err error) {

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", rn.IP), &ssh.ClientConfig{
		User: rn.User,
		Auth: []ssh.AuthMethod{ssh.Password(rn.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			fmt.Printf("Login %s As %s \n", rn.IP, rn.User)
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
