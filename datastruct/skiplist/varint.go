package skiplist

import (
	"encoding/binary"
	"math"
)

func float64Encode(val float64) []byte {
	bits := math.Float64bits(val)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func float64Decode(input []byte) (float64, bool) {
	if len(input) < 8 {
		return 0, false
	}
	bits := binary.LittleEndian.Uint64(input[:8])
	return math.Float64frombits(bits), true
}
