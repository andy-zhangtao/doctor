package node

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"

	zs "github.com/andy-zhangtao/gogather/random"
)

// InitDoctorNode 新增Doctor用户
// 创建用户
// 生成ssh key
// 修改ssh 登录脚本
func InitDoctorNode(name string) (err error) {
	return createUser(name)
}

func createUser(name string) (err error) {
	version := analyLinuxOS()

	logrus.Debugf("OS Version [%s]\n", version)
	switch version {
	case "alpine linux":
		err = alpine(name)
	default:
		// do nothing
		return errors.New("Unsupport Linux OS")
	}

	return
}

func alpine(name string) (err error) {
	// 1. 创建新用户
	var cmd string
	defer func() {
		logrus.Debugf("The exec cmd: [%s]", cmd)
	}()

	cmd = fmt.Sprintf("useradd %s -m -p %s -s /bin/sh", name, zs.GetRandom(8))
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return fmt.Errorf("useradd exec error: %s", err.Error())
	}

	// 2. 创建.ssh目录
	cmd = fmt.Sprintf("mkdir -p /home/%s/.ssh", name)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return fmt.Errorf("create user home error: %s", err.Error())
	}
	// 3. 创建密钥
	cmd = fmt.Sprintf("ssh-keygen -f /home/%s/.ssh/id_rsa -q -N \"\"", name)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return fmt.Errorf("create user key error: %s", err.Error())
	}
	// 4. 创建authorized_keys
	cmd = fmt.Sprintf("mv /home/%s/.ssh/id_rsa.pub /home/%s/.ssh/authorized_keys", name, name)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return fmt.Errorf("create authorized_keys error: %s", err.Error())
	}
	// 5. 修改所属
	cmd = fmt.Sprintf("chown -R %s:%s /home/%s/.ssh/", name, name, name)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return fmt.Errorf("chown user ssh key error: %s", err.Error())
	}
	// 6. 修改权限
	cmd = fmt.Sprintf("chmod 700 /home/%s/.ssh", name)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return fmt.Errorf("chmod .ssh error: %s", err.Error())
	}

	cmd = fmt.Sprintf("chmod 600 /home/%s/.ssh/authorized_keys", name)
	return exec.Command("sh", "-c", cmd).Run()
}

func analyLinuxOS() (linuxOS string) {
	data, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		fmt.Println(err)
		return
	}

	_file := strings.Split(string(data), "\n")

	if len(_file) == 0 {
		return
	}

	name := strings.Split(_file[0], "=")

	if len(name) == 1 {
		return
	}

	_name := strings.ToLower(name[1])

	return _name[1 : len(_name)-1]
}
