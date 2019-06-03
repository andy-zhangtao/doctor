package node

import (
	"bytes"
	"fmt"
	"net"

	"github.com/andy-zhangtao/doctor/doctor/model"
	"golang.org/x/crypto/ssh"
)

// command 在登录SSH后执行特定命令
func command(rn model.RemoteNode, cmd []string) (err error) {
	var conn *ssh.Client
	if rn.Key != "" {
		signer, err := ssh.ParsePrivateKey([]byte(rn.Key))
		if err != nil {
			return fmt.Errorf("unable to parse private key: %v", err)
		}

		conn, err = ssh.Dial("tcp", fmt.Sprintf("%s:22", rn.Ip), &ssh.ClientConfig{
			User: rn.Name,
			Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		})
	} else {
		conn, err = ssh.Dial("tcp", fmt.Sprintf("%s:22", rn.Ip), &ssh.ClientConfig{
			User: rn.Name,
			Auth: []ssh.AuthMethod{ssh.Password(rn.Password)},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		})
	}

	if err != nil {
		return
	}

	session, _ := conn.NewSession()
	defer session.Close()

	for _, c := range cmd {

		var output bytes.Buffer

		session.Stdout = &output
		session.Run(c)

		fmt.Println(output.String())
	}

	return
}
