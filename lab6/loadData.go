package main 

import (
	"fmt"
	"strings"
	"os"
	"encoding/csv"
	"strconv"
	"time"

)



func loadData(filePath string) ([]StockExchangeData, error) {
	if strings.HasSuffix(strings.ToLower(filePath), ".csv") {
		return loadCSV(filePath)
	}
	// if filePath.hasSuffix(".json") {
	// 	return loadJSON(filePath)
	// }
	return nil, fmt.Errorf("unsupported file format: %s", filePath)
}


func loadCSV(filePath string) ([]StockExchangeData, error) {
	file, err := os.Open(filePath)
	errorHandler("Error opening file: ", err)
	defer file.Close()

	reader := csv.NewReader(file)
	
	records, err := reader.ReadAll()
	errorHandler("Error reading CSV file: ", err)

	var data []StockExchangeData
	for i, record := range records {
		if i == 0 {
			continue 
		}
		dateTime, err := time.Parse("01/02/2006", record[0])
		errorHandler("Error parsing date: ", err)

		dateStr := dateTime.Format("2006-01-02")
		


		closeStr := strings.ReplaceAll(record[1], "$", "")
        closeVal, err := strconv.ParseFloat(closeStr, 64)
		if err != nil {
            return nil, fmt.Errorf("error parsing Close value at row %d: %w", i, err)
        }

		// volumeStr := strings.ReplaceAll(record[2], ",", "")
        // volumeVal, err := strconv.Atoi(volumeStr)
        // if err != nil {
        //     return nil, fmt.Errorf("error parsing Volume value at row %d: %w", i, err)
        // }

		openStr := strings.ReplaceAll(record[3], "$", "")
        openVal, err := strconv.ParseFloat(openStr, 64)
        if err != nil {
            return nil, fmt.Errorf("error parsing Open value at row %d: %w", i, err)
        }

		highStr := strings.ReplaceAll(record[4], "$", "")
        highVal, err := strconv.ParseFloat(highStr, 64)
        if err != nil {
            return nil, fmt.Errorf("error parsing High value at row %d: %w", i, err)
        }

		lowStr := strings.ReplaceAll(record[5], "$", "")
        lowVal, err := strconv.ParseFloat(lowStr, 64)
        if err != nil {
            return nil, fmt.Errorf("error parsing Low value at row %d: %w", i, err)
        }

		data = append(data, StockExchangeData{
			Date: dateStr, 
			Close: closeVal,
			Volume: record[2],
			Open: openVal,
			High: highVal,
			Low: lowVal,
		})
	}

	sortDataByDate(data)

	return data, nil
}



