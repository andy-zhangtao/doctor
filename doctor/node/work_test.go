package node

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andy-zhangtao/doctor/model"
)

func TestWork(t *testing.T) {
	rn := model.RemoteNode{
		IP:       "127.0.0.1",
		Password: "123456",
		User:     "test",
	}

	err := Work(rn)
	assert.Nil(t, err)
}
