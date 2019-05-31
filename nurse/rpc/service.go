package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/andy-zhangtao/doctor/nurse/rpc/doctor_v1"
	grpc "google.golang.org/grpc"
)

// ReportNodeInfo 注册节点信息
func ReportNodeInfo(address, ip string, key []byte) (err error) {
	return invoke(address, doctor_v1.DoctorRegister{
		Ip:  ip,
		Key: string(key),
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
