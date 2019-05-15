package node

import (
	"bytes"
	"fmt"
	"net"

	"github.com/andy-zhangtao/doctor/model"
	"golang.org/x/crypto/ssh"
)

// command 在登录SSH后执行特定命令
func command(rn model.RemoteNode, cmd []string) (err error) {
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", rn.IP), &ssh.ClientConfig{
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

	session, _ := conn.NewSession()
	defer session.Close()

	for _, c := range cmd {
		var output bytes.Buffer

		session.Stdout = &output
		session.Run(c)
		fmt.Printf("--- %s ---\n", rn.IP)
		fmt.Println(output.String())
	}

	return
}
