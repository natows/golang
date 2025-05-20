package main 

import (
	"fmt"
)



func (i *ATRIndicator) Name() string {
    return "ATR"
}

func (i *ATRIndicator) Calculate(data []StockExchangeData, day string, days int) ([]float64, error) {
    if len(data) == 0 {
        return nil, fmt.Errorf("no data to calculate ATR")
    }

    index, err := findIndexByDate(data, day)
    if err != nil { 
        return nil, fmt.Errorf("error finding index by date: %w", err)
    }

    data = data[index:]
    

    if len(data) < 2 { 
        return nil, fmt.Errorf("insufficient data to calculate True Range: need at least 2 days, got %d", len(data))
    }

    TRData := []float64{}
    for i := 1; i < len(data); i++ {
        high := data[i].High
        low := data[i].Low
        prevClose := data[i-1].Close
        TRData = append(TRData, max(high-low, abs(high-prevClose), abs(low-prevClose)))
    }

    fmt.Println("True Range Data:", TRData)
    fmt.Println(days)
    if len(TRData) < days { 
        return nil, fmt.Errorf("insufficient True Range values to calculate ATR: need %d, got %d", days, len(TRData))
    }

    ATRData := []float64{}
    firstSum := 0.0
    for i := 0; i < days; i++ {
        firstSum += TRData[i]
    }
    ATRData = append(ATRData, firstSum/float64(days))

    
    for i := 1; i < len(TRData)-days+1; i++ { 
        currentATR := ((float64(days-1) * ATRData[len(ATRData)-1]) + TRData[i+days-1]) / float64(days)
        ATRData = append(ATRData, currentATR)
    }

    return ATRData, nil
}