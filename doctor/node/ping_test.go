package node

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andy-zhangtao/doctor/doctor/model"
)

func TestPing(t *testing.T) {
	rn := model.RemoteNode{
		Ip: "127.0.0.1",
	}

	err := pingNode(rn)

	assert.Nil(t, err)

	rn.Ip = "127.1.1.1"
	err = pingNode(rn)

	assert.NotNil(t, err)
}
