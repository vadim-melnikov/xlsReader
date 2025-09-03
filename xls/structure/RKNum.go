package structure

import (
	"math"
	"strconv"

	"github.com/shakinm/xlsReader/helpers"
)

type RKNum [4]byte

func (r *RKNum) number() (intNum int64, floatNum float64, isFloat bool) {
	rk := helpers.BytesToUint32(r[:])
	isFloat = rk&0x02 == 0
	isMul := rk&0x01 == 1
	if isFloat {
		floatNum = math.Float64frombits(uint64(rk&0xfffffffc) << 32)
		if isMul {
			floatNum /= 100
		}
	} else {
		intNum32 := int32(rk >> 2)
		// Sign extend from 30 bits to 32 bits
		if intNum32&0x20000000 != 0 {
			intNum32 |= ^0x3FFFFFFF // Set upper 2 bits for sign extension
		}
		intNum = int64(intNum32)
		if isMul {
			floatNum = float64(intNum) / 100.
			isFloat = true
		}
	}

	return

}

func (r *RKNum) GetFloat() (fn float64) {
	i, f, isFloat := r.number()
	if isFloat {
		fn = f
	} else {
		fn = float64(i)
	}
	return fn
}

func (r *RKNum) GetInt64() (in int64) {
	i, _, isFloat := r.number()
	if !isFloat {
		in = i
	}
	return in
}

func (r *RKNum) GetString() (s string) {
	i, f, isFloat := r.number()
	if isFloat {
		return strconv.FormatFloat(f, 'f', -1, 64)
	}
	return strconv.FormatInt(i, 10)
}
