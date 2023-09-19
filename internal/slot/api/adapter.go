package api

import (
	"slot-server/internal/slot/api/proto"
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
