package node

import (
	"fmt"
	"os"

	"github.com/andy-zhangtao/doctor/doctor/model"
)

// NodeInit doctor开始工作
// 1. 探测远程节点是否可达
// 2. SFTP向远程节点推送Nurse
// 3. 在远程节点中执行Nurse 初始化
// 	3.1 采集OS信息
// 	3.2 安装Nurse用户
// 	3.3 安装Key
// 4. 初始化结束后，删除Nurse
func NodeInit(rn model.RemoteNode) (err error) {
	err = pingNode(rn)
	if err != nil {
		err = fmt.Errorf("Ping Error: %s", err.Error())
		return
	}

	err = transFileToNode(rn, []string{nurseBIN}, directUp)
	if err != nil {
		err = fmt.Errorf("Sftp Upload Error: %s", err.Error())
		return
	}

	err = command(rn, []string{fmt.Sprintf("%s/d_%s -server %s", destDir, nurseBIN, os.Getenv("DOCTOR_GRPC_SERVER"))})
	if err != nil {
		err = fmt.Errorf("Exec Command In Remote Node Error: %s", err.Error())
		return
	}

	err = command(rn, []string{fmt.Sprintf("rm %s/d_%s", destDir, nurseBIN)})
	if err != nil {
		err = fmt.Errorf("Exec Command In Remote Node Error: %s", err.Error())
		return
	}
	return
}
