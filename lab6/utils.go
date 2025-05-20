package main 

import (
	"sort"
	"fmt"

)



func errorHandler(message string, err error) {
	if err != nil {
		fmt.Println(message, err)
	}

}

func sortDataByDate(data []StockExchangeData) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Date < data[j].Date
	})
}

func findIndexByDate(data []StockExchangeData, day string) (int, error) {
	index := -1
	for i, item := range data {
		if item.Date == (day) {
			return i, nil
		}
	}
	return index, fmt.Errorf("date not found")
}



func max(a, b, c float64) float64 {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}