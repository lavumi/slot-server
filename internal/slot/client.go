package slot

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"slot-server/internal/slot/api/proto"
)

type Client struct {
	proto.SlotClient
}

type AdditionalInfo struct {
	SavedFeature []byte
	CashDiff     float64
	Collectable  bool
}

func Connect() (*Client, error) {
	conn, err := grpc.Dial("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldnt connect to slot Service: %v", err)
		return nil, err
	}

	c := proto.NewSlotClient(conn)
	return &Client{SlotClient: c}, nil
}

func (c *Client) RequestSpin(slotId uint32, bet float32, prevState []byte) ([]byte, *AdditionalInfo, error) {
	req := &proto.Request{
		SlotId:    slotId,
		BetCash:   bet,
		BetLine:   0,
		PrevState: prevState,
	}

	if spin, err := c.Spin(context.Background(), req); err != nil {
		return nil, nil, status.Errorf(codes.Internal, "Error on spin %s", err.Error())
	} else if res, err := protojson.Marshal(spin.GetResult()); err != nil {
		return nil, nil, status.Errorf(codes.DataLoss, "Marshal Spin response failed %s", err.Error())
	} else {
		state := &AdditionalInfo{
			SavedFeature: spin.GetState(),
			CashDiff:     float64(spin.GetCash()),
			Collectable:  false,
		}
		return res, state, nil

	}
}
