package main 

import (
	"fmt"
)

func(i *SMAIndicator) Name() string {
	return "SMA"
}
//sma - srednia arytmetyczna z zamkniec z n dni
func (i *SMAIndicator) Calculate(data []StockExchangeData, day string, days int) ([]float64, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("brak danych do obliczenia SMA")
	}

	index, err := findIndexByDate(data, day)
	errorHandler("Error finding index by date: ", err)

	data = data[index:]

	if len(data) < days {
		return nil, fmt.Errorf("niewystarczające dane do obliczenia SMA dla %d dni", days)
	}
	
	SMAData := []float64{}
	for j:= 0; j <= len(data)-days; j++ {
        sum := 0.0
        for i := 0; i < days; i++ {
			sum += data[i+j].Close

		}
		SMAData = append(SMAData, sum/float64(days))
	}

	return SMAData, nil

} 
