package node

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"

	zs "github.com/andy-zhangtao/gogather/random"
)

// InitDoctorNode 新增Doctor用户
// 创建用户
// 生成ssh key
// 修改ssh 登录脚本
func InitDoctorNode(name string) (key []byte, err error) {
	return createUser(name)
}

func createUser(name string) (key []byte, err error) {
	version := analyLinuxOS()

	logrus.Debugf("OS Version [%s]\n", version)
	switch version {
	case "alpine linux":
		key, err = alpine(name)
	default:
		// do nothing
		return nil, errors.New("Unsupport Linux OS")
	}

	return
}

func alpine(name string) (key []byte, err error) {
	// 1. 创建新用户
	var cmd string
	defer func() {
		logrus.Debugf("The exec cmd: [%s]", cmd)
	}()

	cmd = fmt.Sprintf("useradd %s -m -p %s -s /cli.sh", name, zs.GetRandom(8))
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("useradd exec error: %s", err.Error())
	}

	// 2. 创建.ssh目录
	cmd = fmt.Sprintf("mkdir -p /home/%s/.ssh", name)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("create user home error: %s", err.Error())
	}
	// 3. 创建密钥
	cmd = fmt.Sprintf("ssh-keygen -f /home/%s/.ssh/id_rsa -q -N \"\"", name)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("create user key error: %s", err.Error())
	}
	// 4. 创建authorized_keys
	cmd = fmt.Sprintf("mv /home/%s/.ssh/id_rsa.pub /home/%s/.ssh/authorized_keys", name, name)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("create authorized_keys error: %s", err.Error())
	}
	// 5. 修改所属
	cmd = fmt.Sprintf("chown -R %s:%s /home/%s/.ssh/", name, name, name)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("chown user ssh key error: %s", err.Error())
	}
	// 6. 修改权限
	cmd = fmt.Sprintf("chmod 700 /home/%s/.ssh", name)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("chmod .ssh error: %s", err.Error())
	}

	cmd = fmt.Sprintf("chmod 600 /home/%s/.ssh/authorized_keys", name)
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		return key, fmt.Errorf("chmod authorized_keys error: %s", err.Error())
	}

	if err := copyFile("passwd"); err != nil {
		return key, fmt.Errorf("backup passwd error: %s", err.Error())
	}

	return ioutil.ReadFile(fmt.Sprintf("/home/%s/.ssh/id_rsa", name))
}

func copyFile(filename string) error {
	from, err := os.Open(fmt.Sprintf("/etc/%s", filename))
	if err != nil {
		return err
	}

	defer from.Close()

	to, err := os.OpenFile(fmt.Sprintf("/data/%s", filename), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer to.Close()

	_, err = io.Copy(to, from)
	return err
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
