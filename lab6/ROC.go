package main 

import (
	"fmt"
)


func (i *ROCIndicator) Name() string {
    return "ROC"
}

func (i *ROCIndicator) Calculate(data []StockExchangeData, day string, days int) ([]float64, error) {
	if len(data) < days+1 {
		return nil, fmt.Errorf("not enough data to calculate ROC")
	}

	index, err := findIndexByDate(data, day)
	errorHandler("Error finding index by date: ", err)
	data = data[index:]
	rocData := []float64{}

	for i := days; i < len(data); i++ {
		prev := data[i-days].Close
		if prev == 0 {
			rocData[i-days] = 0 
		} else {
			rocData[i-days] = ((data[i].Close - prev) / prev) * 100
		}
	}

	return rocData, nil
}
