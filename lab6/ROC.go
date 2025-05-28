package main 

import (
	"fmt"
)


func (i *ROCIndicator) Name() string {
    return "ROC"
}

//roc - porownuje cene zamkniecia z cena zamkniecia sprzed n dni
func (i *ROCIndicator) Calculate(data []StockExchangeData, day string, days int) ([]float64, error) {
    if len(data) < days+1 {
        return nil, fmt.Errorf("niewystarczająca liczba danych do obliczenia ROC")
    }

    index, err := findIndexByDate(data, day)
    if err != nil {
        return nil, fmt.Errorf("error finding index by date: %w", err)
    }
    
    if index - days>= 0{
        data = data[index-days:]
    } else {
        return nil, fmt.Errorf("niewystarczająca liczba historycznych danych przed %s: potrzeba conajmniej %d dni", day, days)
    }

    
    if len(data) < days+1 {
        return nil, fmt.Errorf("niewystarczająca ilość danych po podanej dacie: potrzeba conajmniej %d dni, a jest %d", days+1, len(data))
    }
    
    rocData := make([]float64, 0, len(data)-days)
    
    for i := days; i < len(data); i++ {
        current := data[i].Close
        prev := data[i-days].Close
        
        if prev == 0 {
            rocData = append(rocData, 0)
        } else {
            roc := ((current - prev) / prev) * 100
            rocData = append(rocData, roc)
        }
    }
    
    return rocData, nil
}
