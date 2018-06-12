package utils

func CalcAbs(a int64) (ret int64) {
	ret = (a ^ a>>63) - a>>63
	return
}
