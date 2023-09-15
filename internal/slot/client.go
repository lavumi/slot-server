package slot

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"slot-server/internal/server/model"
	"slot-server/internal/slot/api/proto"
)

type Client struct {
	proto.SlotClient
}

func Connect() (*Client, error) {
	conn, err := grpc.Dial("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldnt connect to slot Service: %v", err)
		return nil, err
	}

	c := proto.NewSlotClient(conn).(Client)
	return &c, nil
}

func (c *Client) RequestSpin(slotId uint32, bet float32, prevState string) (*model.SpinOutput, string, error) {
	req := &proto.Request{
		BetCash:   bet,
		BetLine:   0,
		PrevState: prevState,
	}

	if spin, err := c.Spin(context.Background(), req); err != nil {
		return nil, "", err
	} else {

		spinResponse := spin.GetResponse()

		// Slot1 또는 Slot2에 따라 다른 데이터 출력
		switch slotData := spinResponse.(type) {
		case *proto.Response_Foodie:
				slotData.Foodie.
		case *proto.Response_Dragon:

		default:
			return nil, "", status.Errorf(codes.InvalidArgument, "Invalid response type")
		}
	}
}

//func (c *Client) Spin(slotId uint32, bet float32, prevState string) (*model.SpinOutput, string, error) {
//	req := &proto.Request{
//		BetCash:   bet,
//		BetLine:   0,
//		PrevState: prevState,
//	}
//
//	switch slotId {
//	case 0:
//		if spin, err := c.Spin(context.Background(), req); err != nil {
//			return nil, "", err
//		} else {
//			return convertSlotResponse(spin), spin.State, nil
//		}
//	default:
//		return nil, "", fmt.Errorf("invalid SlotId %d", slotId)
//	}
//}
