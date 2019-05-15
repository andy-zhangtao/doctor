package node

import (
	"errors"
	"time"

	"github.com/andy-zhangtao/doctor/model"
	"github.com/sparrc/go-ping"
)

// pingNode 探测远程节点是否可以连通
func pingNode(rn model.RemoteNode) (err error) {
	pinger, err := ping.NewPinger(rn.IP)
	if err != nil {
		return
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		if stats.PacketsRecv == 0 {
			err = errors.New("node unreach! ")
			return
		}
	}

	pinger.Count = 3
	pinger.Timeout = 3 * time.Second
	pinger.Run()
	return
}
