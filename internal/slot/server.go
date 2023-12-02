package slot

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"slot-server/internal/slot/api/proto"
	"slot-server/internal/slot/game/foodie"
)

type Service struct {
	proto.UnimplementedSlotServer
}

func (s *Service) Enter(_ context.Context, req *proto.Request) (*proto.EnterResponse, error) {
	switch req.GetSlotId() {
	case 0:
		if foodieRes, err := foodie.Enter(); err != nil {
			return nil, status.Errorf(codes.Internal, "foodie slot spin error %s", err.Error())
		} else {
			return foodieRes, nil
		}
	default:
		return nil, status.Errorf(codes.Unimplemented, "method Spin not implemented")
	}

}

func (s *Service) Spin(_ context.Context, req *proto.Request) (*proto.SpinResponse, error) {
	switch req.GetSlotId() {
	case 0:
		if foodieRes, err := foodie.Spin(req); err != nil {
			return nil, status.Errorf(codes.Internal, "foodie slot spin error %s", err.Error())
		} else {
			return foodieRes, nil
		}
	default:
		return nil, status.Errorf(codes.Unimplemented, "method Spin not implemented")
	}

}

func Run() {
	slotService := &Service{}

	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("fail to listen port: %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterSlotServer(grpcServer, slotService)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to serve slot Service: %v", err)
	}

	log.Printf("Slot Server Launched\n")
}
