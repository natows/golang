package main 

import (
	"fmt"

)


func main() {
	var indicators = []Indicator{
	&SMAIndicator{},
	&ATRIndicator{},
	&ROCIndicator{},
	}
	
	fmt.Println("Witaj w programie do obliczania wskaźników giełdowych")
	for {
		fmt.Printf("Podaj ścieżkę pliku: [exit by wyjść] ")
		var filePath string
		_, err := fmt.Scanln(&filePath) 
		errorHandler("Błąd w odczycie ścieżki: ", err)
		
		if filePath == "exit" {
            fmt.Println("Do widzenia!")
            break
        }

		// filepath := "HistoricalData_1747726952234.csv"
		dataSource, err := SetDataSource(filePath)
		if err != nil {
            fmt.Println("Błąd, podaj poprawną ścieżkę pliku:", err)
            continue
        }


		data, err := dataSource.LoadData(filePath)
		errorHandler("Błąd przy pobieraniu danych z pliku: ", err)

		date := getDateFromUser(data)
		
		selectedIndicators := getIndicatorsFromUser(indicators)

		
		days := getDaysFromUser(data, date)

		for _, indicator := range selectedIndicators {
			calculateIndicator(data, indicator, date, days)
		}
		
	}
}

