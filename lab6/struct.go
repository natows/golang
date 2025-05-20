package main 





type StockExchangeData struct {
	Date string `json:"date" csv:"date"`
	Close float64 `json:"close" csv:"close"`
	Volume string `json:"volume" csv:"volume"`
	Open float64 `json:"open" csv:"open"`
	High float64 `json:"high" csv:"high"`
	Low float64 `json:"low" csv:"low"`
} 


type Indicator interface {
    Calculate(data []StockExchangeData, startDate string, period int) ([]float64, error)
	Name() string
}

type ATRIndicator struct{}
type SMAIndicator struct{}
type ROCIndicator struct{}






// Date	Data notowania (dzień, w którym odbył się handel na giełdzie).
// Close/Last	Cena zamknięcia – ostatnia cena, po jakiej akcja była handlowana w danym dniu.
// Volume	Wolumen obrotu – liczba akcji, które zostały kupione/sprzedane tego dnia.
// Open	Cena otwarcia – cena, po jakiej rozpoczęto handel daną akcją tego dnia.
// High	Najwyższa cena – najwyższa cena osiągnięta tego dnia.
// Low	Najniższa cena – najniższa cena osiągnięta tego dnia.