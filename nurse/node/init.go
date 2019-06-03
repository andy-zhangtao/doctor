package node

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const user = "nurse"

// InitRemoteNode 新增用户
// 创建用户
// 生成ssh key
// 修改ssh 登录脚本
func InitRemoteNode() (key []byte, err error) {
	key, err = installUser()
	if err != nil {
		return
	}

	return
}

func clearUserInUbuntu() {
	cmd := fmt.Sprintf("rm -rf /home/%s/", user)
	exec.Command("sh", "-c", cmd).Run()
	cmd = fmt.Sprintf("deluser %s", user)
	exec.Command("sh", "-c", cmd).Run()
}

func installUser() (key []byte, err error) {
	return createUser()
}

func createUser() (key []byte, err error) {

	linuxVersion := ""

	switch strings.ToLower(runtime.GOOS) {
	case "linux":
		linuxVersion = getLinuxVersion()
	}

	logrus.Debugf("Node OS Version [%s]", linuxVersion)

	if strings.Contains(linuxVersion, "ubuntu") {
		clearUserInUbuntu()
		if key, err = setupUbuntu(); err != nil {
			return nil, err
		}
		return key, raiseSudo()
	} else if strings.Contains(linuxVersion, "centos") {
		if err := setupCentos(); err != nil {
			return nil, err
		}
	} else if strings.Contains(linuxVersion, "coreos") {
		clearUserInCoreos()
		if key, err = setupCoreos(); err != nil {
			return nil, err
		}
		afterEnableUserInCoreos()
		return key, nil
	}

	return
}

func raiseSudo() (err error) {
	exec.Command("sh", "-c", "mkdir -p /etc/sudoers.d").Run()
	cmd := exec.Command("sh", "-c", "echo \"nurse ALL=(ALL) NOPASSWD: ALL\" > /etc/sudoers.d/nurse")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("raise sudo error [%s]", err)
	}

	if !cmd.ProcessState.Success() {
		return fmt.Errorf("wait raise sudo exec error [%s]", err)
	}

	return
}

func getLinuxVersion() (version string) {
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

func setupUbuntu() (key []byte, err error) {
	cmd := exec.Command("sh", "-c", "useradd -p \"*\" -U -m nurse -s /bin/bash")
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("useradd exec error [%s]", err)
	}

	if !cmd.ProcessState.Success() {
		return nil, fmt.Errorf("wait useradd exec error [%s]", err)
	}

	return ubuntuSetupSSH()
}

func setupCentos() (err error) {
	cmd := exec.Command("sh", "-c", "useradd -p \"*\" -U -m nurse -G wheel")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("useradd exec error [%s]", err)
	}

	if !cmd.ProcessState.Success() {
		return fmt.Errorf("wait useradd exec error [%s]", err)
	}

	return
}
