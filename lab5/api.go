package main 

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func fetchBusStops(url string) []BusStop {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("failed to fetch bus stops:", err)
        return nil
    }
    defer resp.Body.Close()
	var responseData map[string]BusStopData
    err = json.NewDecoder(resp.Body).Decode(&responseData)
    if err != nil {
        fmt.Println("failed to decode JSON:", err)
        return nil
    }

	today := time.Now().Format("2006-01-02")
	dailyData, exists := responseData[today]
	if !exists {
		fmt.Println("No data for today:", today)
		return nil
	}


    busStops := []BusStop{}
	for _,stop := range dailyData.Stops {
		if stop.Name != ""  {
			busStops = append(busStops, stop)
		}
	}
	return busStops

}

func fetchDepartures(url string) []Departure {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("failed to fetch departures:", err)
		return nil
	}
	defer resp.Body.Close()

	var departureData DepartureData
	err = json.NewDecoder(resp.Body).Decode(&departureData) 
	if err != nil {
		fmt.Println("failed to decode JSON:", err)
		return nil
	}
	return departureData.Departures
}

func fetchStopTimes(url string) []StopTimes {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("failed to fetch departures:", err)
		return nil
	}
	defer resp.Body.Close()
	var stopTimesData StopTimesData
	err = json.NewDecoder(resp.Body).Decode(&stopTimesData) 
	if err != nil {
		fmt.Println("failed to decode JSON:", err)
		return nil
	}
	return stopTimesData.StopTimes
}