package models

type CompanyWeeklyMetadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	TimeZone      string `json:"4. Time Zone"`
}

type CompanyDailyMetadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Outputsize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type PriceData struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume int64   `json:"5. volume,string"`
}

type CompanyWeeklyPriceData struct {
	MetaData         CompanyWeeklyMetadata `json:"Meta Data"`
	WeeklyTimeSeries map[string]PriceData  `json:"Weekly Time Series"`
}

type CompanyDailyPriceData struct {
	MetaData         CompanyDailyMetadata `json:"Meta Data"`
	WeeklyTimeSeries map[string]PriceData `json:"Time Series (Daily)"`
}
