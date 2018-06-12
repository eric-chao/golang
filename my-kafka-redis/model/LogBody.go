package model

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
