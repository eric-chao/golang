package test

import (
	"testing"
	. "adhoc/adhoc_data_fast/utils"
)

func Test_Long_Abs(t *testing.T) {

	t.Log(CalcAbs(1528709449))
	t.Log(CalcAbs(-1528709449))
}