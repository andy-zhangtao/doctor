package model

// RemoteNode 远程节点
type RemoteNode struct {
	Ip       string `json:"ip" bw:"ip" pq:"ip"`
	Name     string `pq:"name"`
	Password string `pq:"password"`
	Key      string `pq:"key"`
	Comment  string `pq:"comment"`
}
