package test

import (
	. "adhoc/adhoc_data_fast/utils"
	"testing"
)

func Test_Encoding_String(t *testing.T) {
	key := EncodeKey("prefix", "_history", "a ", "b ", "c ")
	t.Log(key)
}