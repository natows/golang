package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
)



func sortData(data *Response, parameter string) {
	sort.Slice(data.Results, func(i, j int) bool {
		switch parameter {
		case "entity":
			return data.Results[i].Entity < data.Results[j].Entity
		case "currency":
			return data.Results[i].Currency < data.Results[j].Currency
		case "alphabeticcode":
			return data.Results[i].AlphabeticCode < data.Results[j].AlphabeticCode
		case "numericcode":
			i, ierr := strconv.Atoi(data.Results[i].NumericCode)
			j, jerr := strconv.Atoi(data.Results[j].NumericCode)
			if ierr != nil {
				i = 0
			}
			if jerr != nil {
				j = 0
			}
			return i < j
		case "minorunit":
			i, ierr := strconv.Atoi(data.Results[i].MinorUnit)
			j, jerr := strconv.Atoi(data.Results[j].MinorUnit)
			if ierr != nil {
				i = 0
			}
			if jerr != nil {
				j = 0
			}
			return i < j
		default:
			return false	
		}
	})

}

func fetchData() []byte{
	url := "https://public.opendatasoft.com/api/explore/v2.1/catalog/datasets/currency-codes/records?limit=100"
	resp,err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return content
}

type Record struct {
	Entity string `json:"entity"`
	Currency string `json:"currency"`
	AlphabeticCode string `json:"alphabeticcode"`
	NumericCode string `json:"numericcode"`
	MinorUnit string `json:"minorunit"`
}

type Response struct {
	Results []Record `json:"results"`
}


func main() {	
	data := fetchData()
	var response Response
	err := json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	//sortowanie po stringach entity
	sortData(&response, "entity")
	// fmt.Println("Sorted by entity:")
	// for _, record := range response.Results {
	// 	fmt.Printf("Entity: %s\nCurrency: %s\nAlphabetic Code: %s\nNumeric Code: %s\nMinor Unit: %s\n\n", record.Entity, record.Currency, record.AlphabeticCode, record.NumericCode, record.MinorUnit)

	// }

	//sortowanie po intach numeric code
	sortData(&response, "numericcode")
	// fmt.Println("Sorted by numeric code:")
	// for _, record := range response.Results {
	// 	fmt.Printf("Entity: %s\nCurrency: %s\nAlphabetic Code: %s\nNumeric Code: %s\nMinor Unit: %s\n\n", record.Entity, record.Currency, record.AlphabeticCode, record.NumericCode, record.MinorUnit)
	// }

	// //statystyka - liczba unikalnych walut
	currencies := make(map[string]bool)
	for _, record := range response.Results {
		currencies[record.Currency] = true
	}
	fmt.Printf("Number of unique currencies: %d\n", len(currencies))

	// //statystyka - najpopularniejsza waluta
	currencyCount := make(map[string]int)
	for _, record := range response.Results {
		currencyCount[record.AlphabeticCode] +=1
	}
	var maxCurrency []string
	maxCount := 0
	for currency, count := range currencyCount {
		if count > maxCount {
			maxCount = count
			maxCurrency = []string{currency}			
		} else if count == maxCount {
			maxCurrency = append(maxCurrency, currency)
		}
	}
	fmt.Printf("Most popular currencies: %v\n", maxCurrency)


}
