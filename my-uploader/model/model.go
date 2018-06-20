package model

type Result struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Config struct {
	Redis struct {
		Host string `ini:"host"`
		Port string `ini:"port"`
	}

	Storage struct {
		Address string `ini:"address"`
		Path    string `ini:"path"`
		Log     string `ini:"log"`
	}
}
