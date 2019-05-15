package node

import (
	"fmt"

	"github.com/andy-zhangtao/doctor/model"
)

// Work doctor开始工作
// 1. 探测远程节点是否可达
// 2. SFTP向远程节点推送Nurse
// 3. 在远程节点中执行Nurse 初始化
func Work(rn model.RemoteNode) (err error) {
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

	err = command(rn, []string{fmt.Sprintf("%s init", nurseBIN)})
	if err != nil {
		err = fmt.Errorf("Exec Command In Remote Node Error: %s", err.Error())
		return
	}
	return
}
