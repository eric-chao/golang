package test

import (
	. "golang/my-kafka-redis/utils"
	"testing"
)

func Test_Encoding_String(t *testing.T) {
	key := EncodeKey("prefix", "_history", "a ", "b ", "c ")
	t.Log(key)
}
