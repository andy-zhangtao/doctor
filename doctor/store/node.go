package store

import (
	"github.com/andy-zhangtao/doctor/doctor/model"
)

// SaveRemoteNode 保存新主机节点
func SaveRemoteNode(rn model.RemoteNode) (err error) {
	return bwc.Save(rn)
}

func UpdateRemoteNode(rn model.RemoteNode, fields []string) (err error) {
	_, err = bwc.Update(&rn, fields)
	return err
}

func DeleteNode(rn model.RemoteNode, fields []string) (err error) {
	_, err = bwc.Delete(&rn, fields)
	return err
}

// FindSpecNode 查询特定节点数据
func FindSpecNode(rn *model.RemoteNode) (err error) {
	return bwc.FindOne(rn)
}

// FindAllNodes 查询所有节点
func FindAllNodes() ([]*model.RemoteNode, error) {
	var rns []*model.RemoteNode
	err := bwc.FindAllWithSort(&model.RemoteNode{}, &rns, []string{"+ip"})
	return rns, err
}
