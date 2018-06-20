package model

import (
	"bytes"
	"strconv"
	"fmt"
)

type LogBody struct {
	Timestamp  int64
	AppId      string
	ExpId      string
	ModId      string
	ClientId   string
	StatKey    string
	StatValue  float64
	AccValue   float64
	Summary    map[string]string
	Custom     map[string]string
	FromSystem bool
}

func (l LogBody) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("Timestamp: ")
	buf.WriteString(strconv.FormatInt(l.Timestamp, 10))
	buf.WriteString(" | ")

	buf.WriteString("AppId: ")
	buf.WriteString(l.AppId)
	buf.WriteString(" | ")

	buf.WriteString("ExpId: ")
	buf.WriteString(l.ExpId)
	buf.WriteString(" | ")

	buf.WriteString("ModId: ")
	buf.WriteString(l.ModId)
	buf.WriteString(" | ")

	buf.WriteString("ClientId: ")
	buf.WriteString(l.ClientId)
	buf.WriteString(" | ")

	buf.WriteString("StatKey: ")
	buf.WriteString(l.StatKey)
	buf.WriteString(" | ")

	buf.WriteString("StatValue: ")
	buf.WriteString(strconv.FormatFloat(l.StatValue, 'f', 6, 64))
	buf.WriteString(" | ")

	buf.WriteString("AccValue: ")
	buf.WriteString(strconv.FormatFloat(l.AccValue, 'f', 6, 64))
	buf.WriteString(" | ")

	buf.WriteString("Summary: ")
	for k, v := range l.Summary {
		str := fmt.Sprintf("k=%v, v=%v\n ", k, v)
		buf.WriteString(str)
	}
	buf.WriteString(" | ")

	buf.WriteString("Custom: ")
	for k, v := range l.Custom {
		str := fmt.Sprintf("k=%v, v=%v\n ", k, v)
		buf.WriteString(str)
	}
	buf.WriteString(" | ")

	buf.WriteString("FromSystem: ")
	buf.WriteString(strconv.FormatBool(l.FromSystem))

	return buf.String()
}