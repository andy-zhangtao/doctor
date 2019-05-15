package node

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andy-zhangtao/doctor/model"
)

func TestPing(t *testing.T) {
	rn := model.RemoteNode{
		IP: "127.0.0.1",
	}

	err := Ping(rn)

	assert.Nil(t, err)

	rn.IP = "127.1.1.1"
	err = Ping(rn)

	assert.NotNil(t, err)
}
