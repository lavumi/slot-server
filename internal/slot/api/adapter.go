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

func LinePayConvert(wins *[]model.AllLineWin) []*proto.AllLineWin {
	var result []*proto.AllLineWin

	for _, w := range *wins {

		var pos []int32
		for _, strip := range w.Position {
			var bitmask int32
			for _, num := range strip {
				bitmask |= 1 << uint(num)
			}
			pos = append(pos, bitmask)
		}
		win := proto.AllLineWin{
			Win:      float32(w.Win),
			Position: pos,
		}
		result = append(result, &win)
	}
	return result
}
