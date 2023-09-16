package slot

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
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

	c := proto.NewSlotClient(conn)
	return &Client{SlotClient: c}, nil
}

func (c *Client) RequestSpin(slotId uint32, bet float32, prevState string) (*model.SpinOutput, string, error) {
	req := &proto.Request{
		SlotId:    slotId,
		BetCash:   bet,
		BetLine:   0,
		PrevState: prevState,
	}

	if spin, err := c.Spin(context.Background(), req); err != nil {
		return nil, "", err
	} else {
		// protobuf 메시지를 바이트 슬라이스로 직렬화
		res, _ := protojson.Marshal(spin.GetSpinResponse())
		// JSON 문자열 출력
		fmt.Println(string(res))
		fmt.Println(spin.GetState())
	}

	return nil, "", status.Errorf(codes.InvalidArgument, "Invalid response type")
}

func unmarshalProto(baseRes *proto.BaseResult) model.SpinOutput {
	var reel [][]int32
	for _, strip := range baseRes.Reel {
		reel = append(reel, strip.Strip)
	}

	var lw []*model.AllLineWin

	for _, win := range baseRes.LineWins {

		lw = append(lw, &model.AllLineWin{
			Win: win.Win,
			//Position: win.Position,
		})
	}

	return model.SpinOutput{
		Win:         baseRes.Win,
		TotalWin:    baseRes.TotalWin,
		Symbols:     reel,
		UpSymbols:   baseRes.UpSymbol,
		DownSymbols: baseRes.DnSymbol,
		LineWins:    lw,
		BonusWins:   nil,
		NextProcess: nil,
	}
}
