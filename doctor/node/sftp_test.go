package node

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andy-zhangtao/doctor/model"
)

func TestTransFileToNode(t *testing.T) {
	rn := model.RemoteNode{
		IP:       "127.0.0.1",
		Password: "123456",
		User:     "test",
	}

	err := transFileToNode(rn, []string{"TestBin"}, directUp)

	assert.Nil(t, err)
}
