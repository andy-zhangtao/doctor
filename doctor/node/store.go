package node

import (
	"fmt"

	"github.com/andy-zhangtao/doctor/doctor/model"
	"github.com/andy-zhangtao/doctor/doctor/store"
)

func SaveNode(rn model.RemoteNode) error {
	return saveNode(rn)
}

func saveNode(rn model.RemoteNode) (err error) {
	if err = store.SaveRemoteNode(rn); err != nil {
		return fmt.Errorf("Save RemoteNode Error %v", err)
	}

	return
}

func FindOneNode(rn *model.RemoteNode) (err error) {
	return store.FindSpecNode(rn)
}

func FindAllNodes() (rns []model.RemoteNode, err error) {
	rnsPtr, err := store.FindAllNodes()
	if err != nil {
		return rns, err
	}
	for _, r := range rnsPtr {
		rns = append(rns, *r)
	}

	return
}

// UpdateNodeComment 更新节点备注信息
func UpdateNodeComment(rn model.RemoteNode) (err error) {
	return store.UpdateRemoteNode(rn, []string{"ip"})
}

// DeleteNode 删除节点
func DeleteNode(rn model.RemoteNode) (err error) {
	return store.DeleteNode(rn, []string{"ip"})
}
