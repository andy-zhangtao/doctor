package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/andy-zhangtao/doctor/doctor/rpc/doctor_v1"
	"google.golang.org/grpc"
)

// ReportNodeInfo 注册节点信息
func ReportNodeInfo(address string) (err error) {
	return invoke(address, doctor_v1.DoctorRegister{
		Ip: "localhost",
	})
}

func invoke(address string, dr doctor_v1.DoctorRegister) (err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}

	defer conn.Close()

	c := doctor_v1.NewRNRegisterClient(conn)

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	r, err := c.Register(ctx, &dr)
	if err != nil {
		return err
	}

	fmt.Println(r.Msg)
	return
}
