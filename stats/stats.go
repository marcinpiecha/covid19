package stats

import "time"

type Corona struct {
	Confirmed struct {
		Value  int    `json:"value"`
		Detail string `json:"detail"`
	} `json:"confirmed"`
	Recovered struct {
		Value  int    `json:"value"`
		Detail string `json:"detail"`
	} `json:"recovered"`
	Deaths struct {
		Value  int    `json:"value"`
		Detail string `json:"detail"`
	} `json:"deaths"`
	DailySummary    string `json:"dailySummary"`
	DailyTimeSeries struct {
		Pattern string `json:"pattern"`
		Example string `json:"example"`
	} `json:"dailyTimeSeries"`
	Image         string `json:"image"`
	Source        string `json:"source"`
	Countries     string `json:"countries"`
	CountryDetail struct {
		Pattern string `json:"pattern"`
		Example string `json:"example"`
	} `json:"countryDetail"`
	LastUpdate time.Time `json:"lastUpdate"`
}

type Country struct {
	Confirmed struct {
		Value  int    `json:"value"`
		Detail string `json:"detail"`
	} `json:"confirmed"`
	Recovered struct {
		Value  int    `json:"value"`
		Detail string `json:"detail"`
	} `json:"recovered"`
	Deaths struct {
		Value  int    `json:"value"`
		Detail string `json:"detail"`
	} `json:"deaths"`
	LastUpdate time.Time `json:"lastUpdate"`
}

type Url struct {
	Country string
}

//detailed stats

type Global []struct {
	ProvinceState interface{} `json:"provinceState"`
	CountryRegion string      `json:"countryRegion"`
	LastUpdate    int64       `json:"lastUpdate"`
	Confirmed     int         `json:"confirmed"`
	Deaths        int         `json:"deaths"`
	Recovered     int         `json:"recovered"`
	Active        int         `json:"active"`
}
