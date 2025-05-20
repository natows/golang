package main 

import (
	"fmt"

)




func main() {
	var indicators = []Indicator{
	&SMAIndicator{},
	&ATRIndicator{},
	// &ROCIndicator{},
	}

	fmt.Printf("Podaj ścieżkę pliku: ")
	var filePath string
    _, err := fmt.Scanln(&filePath) 
	errorHandler("Error reading file path: ", err)

	// filepath := "HistoricalData_1747726952234.csv"
	data, err := loadData(filePath)
	
	errorHandler("Error loading data: ", err)

	fmt.Println("Loaded data:", data)
	var date string
	fmt.Printf("Podaj datę (format: yyyy-mm-dd): ")
	for {
		_, err = fmt.Scanln(&date)
		errorHandler("Error reading date: ", err)
		_, err := findIndexByDate(data, date)
		if err != nil {
			fmt.Printf("Nie znaleziono daty %s w danych. Podaj ponownie: ", date)
		} else {
			break
		}
	}
	
	var days int
	fmt.Printf("Podaj liczbę dni: ")
	for {
		_, err = fmt.Scanln(&days)
		errorHandler("Error reading days: ", err)
		index, err := findIndexByDate(data, date)
		errorHandler("Error finding index by date: ", err)
		if len(data[index:]) < days {
		fmt.Printf("Nie można obliczyć wskaźników, ponieważ nie ma wystarczającej liczby danych. Podaj mniejsza liczbę: ")
		}else {
			break
		}
	}

	for _, indicator := range indicators {
		calculateIndicator(data, indicator, date, days)
	}


}


func calculateIndicator(data []StockExchangeData, indicator Indicator, day string, days int) {
    results, err := indicator.Calculate(data, day, days)
    if err != nil {
        fmt.Println("Error calculating indicator:", err)
        return
    }
    index, err := findIndexByDate(data, day)
    errorHandler("Error finding index by date: ", err) 

    fmt.Printf("%s \n", indicator.Name())
    for i := 0; i < len(results); i++ {
        if index+i+days-1 < len(data) {
            fmt.Printf("%s - %s: %f \n", data[index+i].Date, data[index+i+days-1].Date, results[i])
        } 
    }
    fmt.Println("--------------------------------------------------")

}