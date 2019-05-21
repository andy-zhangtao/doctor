package impl

import (
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"

	"github.com/andy-zhangtao/doctor/doctor/rpc/doctor_v1"
)

type server struct{}

func (s *server) Register(ctx context.Context, in *doctor_v1.DoctorRegister) (*doctor_v1.Reply, error) {

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
