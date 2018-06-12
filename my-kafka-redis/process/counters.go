package process

import (
	. "adhoc/adhoc_data_fast/model"
	. "adhoc/adhoc_data_fast/config"
	. "adhoc/adhoc_data_fast/utils"
)

var AllCounter = &NewLog{
	Prefix:      "adhoc_all",
	CustomNames: []string{"experiment_id"},
	CustomKeys: func(body LogBody) []string {
		return []string{body.ExpId}
	},
	TimeString: func(t int64) string {
		return "all"
	},
	IgnoreDays:    GlobalConfig.Adhoc.IgnoreDays,
	ExpireDays:    GlobalConfig.Adhoc.ExpireDays,
	CountStat:     true,
	CountVariance: true,
	UseLogLog:     false,
}

var HourlyCounter = &NewLog{
	Prefix:      "adhoc_hourly",
	CustomNames: []string{"experiment_id"},
	CustomKeys: func(body LogBody) []string {
		return []string{body.ExpId}
	},
	TimeString: func(t int64) string {
		return Hourly(t)
	},
	IgnoreDays:    GlobalConfig.Adhoc.IgnoreDays,
	ExpireDays:    GlobalConfig.Adhoc.IgnoreDays + 1,
	CountStat:     true,
	CountVariance: false,
	UseLogLog:     true,
}

var DailyCounter = &NewLog{
	Prefix:      "adhoc_daily",
	CustomNames: []string{"experiment_id"},
	CustomKeys: func(body LogBody) []string {
		return []string{body.ExpId}
	},
	TimeString: func(t int64) string {
		return Daily(t)
	},
	IgnoreDays:    GlobalConfig.Adhoc.IgnoreDays,
	ExpireDays:    GlobalConfig.Adhoc.IgnoreDays + 1,
	CountStat:     true,
	CountVariance: false,
	UseLogLog:     true,
}

var MonthlyUvCounter = &NewLog{
	Prefix:      "adhoc_monthly_uv",
	CustomNames: []string{},
	CustomKeys: func(body LogBody) []string {
		return []string{}
	},
	TimeString: func(t int64) string {
		return Monthly(t)
	},
	IgnoreDays:    GlobalConfig.Adhoc.IgnoreDays,
	ExpireDays:    32,
	CountStat:     true,
	CountVariance: false,
	UseLogLog:     true,
}
