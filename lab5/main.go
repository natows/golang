package main

import (
	"fmt"
	"strings"
	"sync"
)

func printDepartures(departures []Departure, chosenStop BusStop, display int) {
	fmt.Println("\n=== Odjazdy z przystanku", chosenStop, "===")
	fmt.Printf("%-5s %-8s %-25s %-16s %-12s\n", "Lp.", "Linia", "Kierunek", "Planowany odjazd", "Opóźnienie")
	fmt.Println(strings.Repeat("-", 70))

	displayCount := display
	if len(departures) < 10 {
		displayCount = len(departures)
	}

	for i := 0; i < displayCount; i++ {
		dep := departures[i]

		departureTime := formatTimeToLocal(dep.Estimated)
		
		direction := dep.Headsign
	
		delayText := "punktualnie"
		if dep.Delay > 0 {
			delayText = fmt.Sprintf("+%ds", dep.Delay)
		} else if dep.Delay < 0 {
			delayText = fmt.Sprintf("-%ds", -dep.Delay)
		}
		
		fmt.Printf("%-5d %-8d %-25s %-16s %-12s\n", 
				i+1, dep.RouteID, direction, departureTime, delayText)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	url := "https://ckan.multimediagdansk.pl/dataset/c24aa637-3619-4dc2-a171-a23eec8f2172/resource/4c4025f0-01bf-41f7-a39f-d156d201b82b/download/stops.json"
	busStops := fetchBusStops(url)
	if busStops == nil {
		fmt.Println("No bus stops found")
		return
	}

	for {
		fmt.Printf("=========================\n")
		chosenStop := getStopFromUser(busStops)
		fmt.Printf("Wybrany przystanek: %s %s\n", chosenStop.Name, chosenStop.StopCode)
		fmt.Printf("\n========================\n")
		fmt.Printf("1. Monitoruj jedną linię\n")
		fmt.Printf("2. Monitoruj dwie linie\n")
		fmt.Printf("3. Wyjdź\n")
		fmt.Printf("========================\n")
		fmt.Printf("Wybierz opcję: ")
		var optionChoice int
        _, err := fmt.Scanln(&optionChoice) 
        if err != nil {
            fmt.Println("Błąd odczytu. Spróbuj ponownie.")
            var discard string
            fmt.Scanln(&discard)
            continue
        }
		baseURL := "https://ckan2.multimediagdansk.pl/departures"
		fullURL := fmt.Sprintf("%s?stopId=%d", baseURL, chosenStop.ID)
		departures := fetchDepartures(fullURL)
		if len(departures) == 0  {
			fmt.Println("Nie znaleziono odjazdów dla przystanku", chosenStop.Name)
			continue
		}

		displayCount := 10
		switch optionChoice {
			case 1: 
				wg.Add(1)
				printDepartures(departures, chosenStop, displayCount)
				monitorLine(wg, mutex, busStops, chosenStop, 1, displayCount, departures)
			case 2: 
				wg.Add(2)
				printDepartures(departures, chosenStop, displayCount)
				go monitorLine(wg, mutex, busStops, chosenStop, 1, displayCount, departures)
				go monitorLine(wg, mutex, busStops, chosenStop, 2, displayCount, departures)
				wg.Wait()
			case 3: 
				fmt.Println("Do widzenia!")
				return
			default:
				fmt.Println("Nieprawidłowy wybór")
				continue
			}

	}


}