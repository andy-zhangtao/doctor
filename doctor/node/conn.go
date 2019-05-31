package node

var nodeRegisterChanMap map[string]chan string

// InitConnChan 初始化通道Map
// key放置的是远程节点IP
func InitConnChan() {
	if nodeRegisterChanMap == nil {
		nodeRegisterChanMap = make(map[string]chan string)
	}

	return
}

// GetChan 获取相对应IP的chan
func GetChan(key string) chan string {
	return nodeRegisterChanMap[key]
}

// SetupChan 生成新的chan
func SetupChan(key string) {
	if _, ok := nodeRegisterChanMap[key]; !ok {
		c := make(chan string)
		nodeRegisterChanMap[key] = c
	}

	return
}
