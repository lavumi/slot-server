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
		strip := proto.Strip{Symbol: convertIntArray[int32](s)}
		result = append(result, &strip)
	}
	return result
}

func LinePayConvert(wins *[]model.AllLineWin) []*proto.AllLineWin {
	var result []*proto.AllLineWin
	for _, w := range *wins {
		win := proto.AllLineWin{
			Win:      float32(w.Win),
			Position: nil,
		}
		result = append(result, &win)
	}
	return result
}
