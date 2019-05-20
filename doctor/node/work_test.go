package node

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andy-zhangtao/doctor/doctor/model"
)

func TestWork(t *testing.T) {
	rn := model.RemoteNode{
		Ip:       "127.0.0.1",
		Password: "123456",
		Name:     "test",
	}

	err := NodeInit(rn)
	assert.Nil(t, err)
}
