package main 

import (
	"fmt"
)



func (i *ATRIndicator) Name() string {
    return "ATR"
}

//atr - srednia z maksymalnych roznic miedzy cenami zamkniec, otwarc, wysokich i niskich z n dni

func (i *ATRIndicator) Calculate(data []StockExchangeData, day string, days int) ([]float64, error) {
    if len(data) == 0 {
        return nil, fmt.Errorf("brak danych do obliczenia ATR")
    }

    index, err := findIndexByDate(data, day)
    errorHandler("Error finding index by date: ", err)

    if index < days {
        return nil, fmt.Errorf("niewystarczająca ilość danych przed %s do obliczenia ATR(%d): potrzeba %d dni przed", day, days, days)
    }

    data = data[index-days:]

    if len(data) < days+1 { 
        return nil, fmt.Errorf("insufficient data to calculate ATR: need at least %d days", days+1)
    }

    TRData := make([]float64, len(data)-1)
    for i := 1; i < len(data); i++ {
        high := data[i].High
        low := data[i].Low
        prevClose := data[i-1].Close
        TRData[i-1] = max(high-low, abs(high-prevClose), abs(low-prevClose))
    }

    if len(TRData) < days { 
        return nil, fmt.Errorf("insufficient True Range values to calculate ATR: need %d, got %d", days, len(TRData))
    }

    resultLen := len(TRData) - days + 1
    ATRData := make([]float64, 0, resultLen)

    firstSum := 0.0
    for i := 0; i < days; i++ {
        firstSum += TRData[i]
    }
    ATRData = append(ATRData, firstSum/float64(days))

    for i := 1; i < resultLen; i++ {
        currentATR := ((float64(days-1) * ATRData[i-1]) + TRData[i+days-1]) / float64(days)
        ATRData = append(ATRData, currentATR)
    }

    return ATRData, nil
}