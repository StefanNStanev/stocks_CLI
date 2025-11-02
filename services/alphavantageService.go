package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/StefanNSTanev/stocks_CLI/models"
)

var API_KEY = "AT5H0AX72LAQQAEO"

func GetWeeklyData(symbol string) (models.CompanyWeeklyPriceData, error) {
	var weeklyData models.CompanyWeeklyPriceData
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=%s&apikey=%s", symbol, API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		return weeklyData, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return weeklyData, err
	}

	err = json.Unmarshal(body, &weeklyData)
	if err != nil {
		return weeklyData, err
	}

	return weeklyData, nil
}

func GetDailyData(symbol string) (models.CompanyDailyPriceData, error) {
	var dailyData models.CompanyDailyPriceData
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", symbol, API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		return dailyData, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dailyData, err
	}

	err = json.Unmarshal(body, &dailyData)
	if err != nil {
		return dailyData, err
	}

	return dailyData, nil
}
