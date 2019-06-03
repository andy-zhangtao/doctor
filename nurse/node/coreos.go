package node

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

type coreUser struct {
	User string
	Key  string
}

const cloudConfig = `#cloud-config
users:
  - name: {{.User}}
    ssh_authorized_keys:
      - {{.Key}}
    groups:
      - sudo
      - docker
    shell: /bin/sh
write_files:
  - path: /etc/ssh/sshd_config
    permissions: 0600
    owner: root
    content: |
      # Use most defaults for sshd configuration.
      UsePrivilegeSeparation sandbox
      Subsystem sftp internal-sftp
      ClientAliveInterval 180
      UseDNS no
      PermitRootLogin yes
      ChallengeResponseAuthentication no
      PasswordAuthentication yes`

func setupCoreos() (key []byte, err error) {
	private, public, err := coreosSetupSSH()
	if err != nil {
		return
	}

	c := coreUser{
		User: user,
		Key:  string(public),
	}

	t := template.Must(template.New("config").Parse(cloudConfig))

	f, err := os.Create("/tmp/nurse_cloud_config")
	defer f.Close()

	if err != nil {
		return key, fmt.Errorf("create config error: %s", err)
	}
	if err = t.Execute(f, c); err != nil {
		return key, fmt.Errorf("create nuse config error: %s", err)
	}

	if err = enableUserInCoreos(); err != nil {
		return key, fmt.Errorf("enable cloud_config error: %s", err.Error())
	}

	return private, nil
}

func enableUserInCoreos() error {
	cmd := fmt.Sprintf("coreos-cloudinit --from-file=/tmp/%s_cloud_config", user)
	return exec.Command("sh", "-c", cmd).Run()
}

func clearUserInCoreos() {
	cmd := fmt.Sprintf("userdel -f -r %s", user)
	exec.Command("sh", "-c", cmd).Run()

	cmd = fmt.Sprintf("rm -rf /tmp/%s_*", user)
	exec.Command("sh", "-c", cmd).Run()
	return
}

func afterEnableUserInCoreos() {
	cmd := fmt.Sprintf("rm -rf /tmp/%s_*", user)
	exec.Command("sh", "-c", cmd).Run()
	return
}
