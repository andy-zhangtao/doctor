package impl

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"

	"github.com/andy-zhangtao/doctor/doctor/node"
	"github.com/andy-zhangtao/doctor/doctor/rpc/doctor_v1"
)

type server struct{}

func (s *server) Register(ctx context.Context, in *doctor_v1.DoctorRegister) (*doctor_v1.Reply, error) {

	// 接受到远程节点注册信息
	// rn := model.RemoteNode{
	// 	Ip:  in.Ip,
	// 	Key: in.Key,
	// }

	// if err := store.SaveRemoteNode(rn); err != nil {
	// 	logrus.Errorf("Node Register Error: [%s]", err.Error())
	// 	return nil, err
	// }

	logrus.Debugf("Docter Got Nurse [%s] Register Message ", in.Ip)
	node.GetChan(in.Ip) <- in.Key

	logrus.Debugf("Send [%s] Register Message ", in.Ip)
	return &doctor_v1.Reply{
		Msg: "OK",
	}, nil
}

// Run 启动GRPC服务
func Run() (err error) {
	port := "50000"
	if os.Getenv("DOCTOR_GRPC_PORT") != "" {
		port = os.Getenv("DOCTOR_GRPC_PORT")
	}

	ser, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return fmt.Errorf("Create GRPC Server Error %s", err.Error())
	}

	s := grpc.NewServer()

	doctor_v1.RegisterRNRegisterServer(s, &server{})
	reflection.Register(s)
	return s.Serve(ser)
}
