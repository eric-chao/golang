package model

/*
{
  "app_key": "ADHOC_d1a5fa92-d874-48d6-bc9b-46056563e671",
  "client_id": "869014027066125",
  "stats": [
    {
    "key": "t1",
    "value": 1,
    "acc_value": 1,
    "timestamp": 1528416000,
	"experiment_ids": [
		"3b5d0e29-47ae-4123-a9cf-05313ac5caba"
	]
    }
  ],
  "summary": {
    "device_name": "Xiaomi Mi4",
    "device_id": "b6005083cfbaad22",
    "display_height": 1920,
    "os_version": "5.1",
    "device_model": "Xiaomi Mi4",
    "app_version": "4.3.3.0.31",
    "display_width": 1080,
    "country": "CN",
    "sdk_api_version": 2.1,
    "locale": "zh_CN",
    "os_version_name": "5.1",
    "wifi_mac": "68:3e:34:c3:d7:0d",
    "language": "zh",
    "package_name": "cn.yonghui.hyd",
    "sdk_version": "3.1.5",
    "network_state": "WIFI_CONNECTED",
    "region_province": "中国大陆",
    "OS": "google_android",
    "region_city": "安徽省",
    "screen_size": 2
  },
  "custom": {}
}
 */
type Stat struct {
	Key           string   `json:"key"`
	Value         float64  `json:"value"`
	Timestamp     int64    `json:"timestamp"`
	AccValue      float64  `json:"acc_value"`
	ExperimentIds []string `json:"experiment_ids"`
}

type RequestBody struct {
	AppKey   string            `json:"app_key"`
	ClientId string            `json:"client_id"`
	Stats    []Stat            `json:"stats"`
	summary  map[string]string `json:"summary"`
	custom   map[string]string `json:"custom"`
}
