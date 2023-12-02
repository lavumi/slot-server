package api

import (
	"slot-server/internal/slot/api/proto"
	"slot-server/internal/slot/model"
)

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func convertIntArray[U, T Int](s []T) (out []U) {
	out = make([]U, len(s))
	for i := range s {
		out[i] = U(s[i])
	}
	return out
}

func GridConvert(grid [][]int) []*proto.Strip {

	var result []*proto.Strip

	for _, s := range grid {
		strip := proto.Strip{Strip: convertIntArray[int32](s)}
		result = append(result, &strip)
	}
	return result
}

func PayLineConvert(pays []model.Line) []*proto.Payout {
	var result []*proto.Payout

	for _, p := range pays {
		payout := proto.Payout{
			Symbol: uint32(p.Symbol),
			Payout: p.Payout,
		}
		result = append(result, &payout)
	}

	return result
}

func SymbolTableConvert(symbols map[string]string) []*proto.Symbol {
	var result []*proto.Symbol

	for key, value := range symbols {
		payout := proto.Symbol{
			Index: key,
			Img:   value,
		}
		result = append(result, &payout)
	}

	return result
}
