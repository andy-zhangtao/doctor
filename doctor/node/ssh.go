package node

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/andy-zhangtao/doctor/doctor/model"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func Login(rn model.RemoteNode) error {
	return sshRemoteNode(rn)
}

func sshRemoteNode(rn model.RemoteNode) (err error) {
	signer, err := ssh.ParsePrivateKey([]byte(rn.Key))
	if err != nil {
		return fmt.Errorf("unable to parse private key: %v", err)
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", rn.Ip), &ssh.ClientConfig{
		User: rn.Name,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			fmt.Printf("Login %s As %s Use Private Key \n", rn.Ip, rn.Name)
			return nil
		},
	})

	if err != nil {
		return fmt.Errorf("login %s error: %v", rn.Ip, err)
	}

	session, err := client.NewSession()
	defer session.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	go func(session *ssh.Session) {
		for {
			select {
			case <-c:
				if err = session.Signal(ssh.SIGKILL); err != nil {
					fmt.Println(err)
				}
			}
		}
	}(session)

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	termFD := int(os.Stdin.Fd())
	w, h, _ := terminal.GetSize(termFD)
	termState, _ := terminal.MakeRaw(termFD)

	defer func() {
		signal.Stop(c)
		session.Close()
		terminal.Restore(termFD, termState)
	}()

	err = session.RequestPty("xterm", h, w, modes)
	if err != nil {
		return
	}
	err = session.Shell()
	if err != nil {
		return
	}

	return session.Wait()

}
