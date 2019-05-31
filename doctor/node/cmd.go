package node

import (
	"bytes"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"

	"github.com/andy-zhangtao/doctor/doctor/model"
	"golang.org/x/crypto/ssh"
)

// command 在登录SSH后执行特定命令
func command(rn model.RemoteNode, cmd []string) (err error) {
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", rn.Ip), &ssh.ClientConfig{
		User: rn.Name,
		Auth: []ssh.AuthMethod{ssh.Password(rn.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			logrus.Debugf("Login %s As %s \n", rn.Ip, rn.Name)
			return nil
		},
	})

	if err != nil {
		return
	}

	session, _ := conn.NewSession()
	defer session.Close()

	for _, c := range cmd {
		fmt.Println(c)
		var output bytes.Buffer

		session.Stdout = &output
		session.Run(c)
		fmt.Printf("--- %s ---\n", rn.Ip)
		fmt.Println(output.String())
	}

	return
}
