package main 

import (
	"sort"
	"fmt"
	"strings"

)



func errorHandler(message string, err error) {
	if err != nil {
		fmt.Println(message, err)
		return
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


func calculateIndicator(data []StockExchangeData, indicator Indicator, day string, days int) {
    results, err := indicator.Calculate(data, day, days)
    errorHandler("Error calculating indicator: ", err)
    index, err := findIndexByDate(data, day)
    errorHandler("Error finding index by date: ", err) 

    fmt.Printf("%s \n", indicator.Name())
	printingLen := 10
	if len(results) < printingLen {
		printingLen = len(results)
	}
    for i := 0; i < printingLen; i++ {
        if index+i+days-1 < len(data) {
            fmt.Printf("%s - %s: %f \n", data[index+i].Date, data[index+i+days-1].Date, results[i])
        } 
    }
	if printingLen < len(results) {
		fmt.Printf("...\n")
	}
    fmt.Println("--------------------------------------------------")

}
func getDateFromUser(data []StockExchangeData) string {
	var date string
	fmt.Printf("Dostępne daty: [%s - %s]\n", data[0].Date, data[len(data)-1].Date)
	fmt.Printf("Podaj datę (format: yyyy-mm-dd): ")
	for {
		_, err := fmt.Scanln(&date)
		errorHandler("Error reading date: ", err)
		_, err = findIndexByDate(data, date)
		if err != nil {
			fmt.Printf("Nie znaleziono daty %s w danych. Podaj ponownie: ", date)
		} else {
			break
		}
	}
	return date	
}

func getIndicatorsFromUser(indicators []Indicator) []Indicator {
	fmt.Printf("Które wskaźniki chcesz obliczyć? (podaj liczby po przecinku)(1 - SMA, 2 - ATR, 3 - ROC, 4 - wszystkie): ")
	var choice string
	_, err := fmt.Scanln(&choice)
	errorHandler("Error reading choice: ", err)
	split := strings.Split(choice, ",")
	var selectedIndicators []Indicator
	for _, s := range split {
		switch s {
		case "1":
			selectedIndicators = append(selectedIndicators, &SMAIndicator{})
		case "2":
			selectedIndicators = append(selectedIndicators, &ATRIndicator{})
		case "3":
			selectedIndicators = append(selectedIndicators, &ROCIndicator{})
		case "4":
			selectedIndicators = indicators
		default:
			fmt.Println("Nieprawidłowy wybór. Wybierz ponownie.")
			continue
		}
	}
	return selectedIndicators
}
func getDaysFromUser(data []StockExchangeData, date string) int {
	var days int
	fmt.Printf("Podaj liczbę dni: ")
	for {
		_, err := fmt.Scanln(&days)
		errorHandler("Error reading days: ", err)
		index, err := findIndexByDate(data, date)
		errorHandler("Error finding index by date: ", err)
		if len(data[index:]) < days {
			fmt.Printf("Nie można obliczyć wskaźników, ponieważ nie ma wystarczającej liczby danych. Podaj mniejsza liczbę dni: ")
		}else {
			break
		}
	}
	return days
}