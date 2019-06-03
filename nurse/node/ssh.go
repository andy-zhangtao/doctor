package node

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func ubuntuSetupSSH() (key []byte, err error) {
	var cmd string
	defer func() {
		logrus.Debugf("The exec cmd: [%s]", cmd)
	}()

	// 2. 创建.ssh目录
	cmd = fmt.Sprintf("mkdir -p /home/%s/.ssh", user)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("create user home error: %s", err.Error())
	}
	// 3. 创建密钥
	cmd = fmt.Sprintf("ssh-keygen -f /home/%s/.ssh/id_rsa -q -N \"\"", user)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("create user key error: %s", err.Error())
	}
	// 4. 创建authorized_keys
	cmd = fmt.Sprintf("mv /home/%s/.ssh/id_rsa.pub /home/%s/.ssh/authorized_keys", user, user)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("create authorized_keys error: %s", err.Error())
	}
	// 5. 修改所属
	cmd = fmt.Sprintf("chown -R %s:%s /home/%s/.ssh/", user, user, user)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("chown user ssh key error: %s", err.Error())
	}
	// 6. 修改权限
	cmd = fmt.Sprintf("chmod 700 /home/%s/.ssh", user)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("chmod .ssh error: %s", err.Error())
	}

	cmd = fmt.Sprintf("chmod 600 /home/%s/.ssh/authorized_keys", user)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("chmod authorized_keys error: %s", err.Error())
	}

	return ioutil.ReadFile(fmt.Sprintf("/home/%s/.ssh/id_rsa", user))
}

func coreosSetupSSH() (privateKey, publicKey []byte, err error) {
	cmd := fmt.Sprintf("ssh-keygen -f /tmp/%s_id_rsa -q -N \"\"", user)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return privateKey, publicKey, fmt.Errorf("create user key error: %s", err.Error())
	}

	privateKey, err = ioutil.ReadFile(fmt.Sprintf("/tmp/%s_id_rsa", user))
	if err != nil {
		return
	}

	publicKey, err = ioutil.ReadFile(fmt.Sprintf("/tmp/%s_id_rsa.pub", user))
	if err != nil {
		return
	}

	return
}
