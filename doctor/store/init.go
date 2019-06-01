package store

import (
	"fmt"
	"os"

	"github.com/andy-zhangtao/bwidow"
	"github.com/andy-zhangtao/doctor/doctor/model"
)

var bwc *bwidow.BW

func init() {
	bwc = bwidow.GetWidow().Driver(bwidow.DRIVER_PQ)
	if bwc.Error() != nil {
		fmt.Printf("BW Init Error: %v \n", bwc.Error())
		os.Exit(-1)
	}

	fmt.Printf("BW Version %v \n", bwc.Version())

	bwc.Map(model.RemoteNode{}, "doctor_remote_node")
	return
}
