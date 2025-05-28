package main 

import (
	"fmt"
	"sync"
	"time"
)


func chooseDeparture(mutex *sync.Mutex, departures []Departure, workerNumber int, displayCount int) (Departure, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var choice int
	fmt.Printf("\n[WORKER %d] Wybierz numer odjazdu (1-" + fmt.Sprintf("%d", displayCount) + "): ", workerNumber)
	_, err := fmt.Scanln(&choice)

	if err != nil || choice < 1 || choice > displayCount {
		return Departure{}, fmt.Errorf("Nieprawidłowy wybór")
	}


	chosenDeparture := departures[choice-1]	
	return chosenDeparture, nil
}

func printMonitoringData(mutex *sync.Mutex, workerNumber int, chosenDeparture Departure, chosenStop BusStop, chosenStopTime StopTimes, nextStop BusStop, secondStop BusStop, end1 bool, end2 bool, arrival1 string, arrival2 string, dep1 Departure, dep2 Departure) {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("\n[WORKER %d] Wybrałeś linię %d w kierunku %s na przystanku %s %s\n", 
	workerNumber, chosenDeparture.RouteID, chosenDeparture.Headsign, chosenStop.Name, chosenStop.StopCode)
	fmt.Printf("Planowany przyjazd na przystanek %s %s: %s.\n %s\n", chosenStop.Name, chosenStop.StopCode, formatTimeToLocal(chosenStopTime.ArrivalTime), formatDelayTime(chosenDeparture.Delay))
	if end1 {
		fmt.Printf("Koniec trasy\n")
	}
	fmt.Printf("Planowany przyjazd na przystanek %s %s: %s.\n %s\n", nextStop.Name, nextStop.StopCode, arrival1, formatDelayTime(dep1.Delay))
	if end2 {
		fmt.Printf("Koniec trasy\n")
	}
	fmt.Printf("Planowany przyjazd na przystanek %s %s: %s.\n %s\n\n", secondStop.Name, secondStop.StopCode, arrival2, formatDelayTime(dep2.Delay))

}

func monitorLine(wg *sync.WaitGroup, mutex *sync.Mutex, busStops []BusStop, chosenStop BusStop, workerNumber int, displayCount int, departures []Departure) {
	defer wg.Done()

	chosenDeparture, err := chooseDeparture(mutex, departures, workerNumber, displayCount)
	if err != nil {
		fmt.Printf("[WORKER %d] Błąd: %s", workerNumber, err)
		return
	}


	timestamp := time.Now().Format("2006-01-02")
	url := fmt.Sprintf("https://ckan2.multimediagdansk.pl/stopTimes?date=%s&routeId=%d",timestamp, chosenDeparture.RouteID)
	stopTimes := fetchStopTimes(url)


	theoricalTime := formatTimeToLocal(chosenDeparture.Theoretical)
 
	index := findStopTime(stopTimes, chosenDeparture, chosenStop, theoricalTime)

	if index == -1 {
		fmt.Printf("[WORKER %d] Nie znaleziono przystanku w rozkładzie!", workerNumber)
		return
	}
	

	chosenStopTime := stopTimes[index]
	nextStopTime:= stopTimes[index+1]
	secondStopTime := stopTimes[index+2]

	var nextStop BusStop
	var secondStop BusStop

	end1 := false

	if checkIfEnd(nextStopTime) {
		end1 = true
	} else {
		nextStop, err = findStop(busStops, nextStopTime.StopId)
		if err != nil {
			fmt.Println("Błąd:", err)
			return
		}
	}

	end2 := false
	if checkIfEnd(secondStopTime) {
		end2 = true
	}else {
		secondStop, err = findStop(busStops, secondStopTime.StopId)
		if err != nil {
			fmt.Println("Błąd:", err)
			return
		}

	}
	arrival1 := formatTimeToLocal(nextStopTime.ArrivalTime)

	arrival2 := formatTimeToLocal(secondStopTime.ArrivalTime)

	departures1 := fetchDepartures(fmt.Sprintf("https://ckan2.multimediagdansk.pl/departures?stopId=%d", nextStopTime.StopId))
	departures2 := fetchDepartures(fmt.Sprintf("https://ckan2.multimediagdansk.pl/departures?stopId=%d", secondStopTime.StopId))
	
	var dep1, dep2 Departure
	dep1, err = findDeparture(departures1, chosenDeparture.RouteID, formatTimeToLocal(nextStopTime.DepartureTime))
	if err != nil {
		fmt.Println("Błąd:", err)
		return
	}


	dep2, err = findDeparture(departures2, chosenDeparture.RouteID, formatTimeToLocal(secondStopTime.DepartureTime))
	if err != nil {
		fmt.Println("Błąd:", err)
		return
	}
	printMonitoringData(mutex, workerNumber, chosenDeparture, chosenStop, chosenStopTime, nextStop, secondStop, end1, end2, arrival1, arrival2, dep1, dep2)
}