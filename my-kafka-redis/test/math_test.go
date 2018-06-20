package test

import (
	"testing"
	. "adhoc/adhoc_data_fast_golang/utils"
)

func Test_Long_Abs(t *testing.T) {

	t.Log(CalcAbs(1528709449))
	t.Log(CalcAbs(-1528709449))
}

func Test_Long_Equals(t *testing.T) {
	var l int64 = 1
	var m int32 = 1
	var n int16 = 1
	var s int8 = 1

	t.Log(l == 1, m == 1, n == 1, s == 1)
}
