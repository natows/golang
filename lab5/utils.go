package main 

import (
	"fmt"
	"strings"
	"time"

)
func getStopFromUser(busStops []BusStop) BusStop {
	fmt.Print("Wpisz nazwę lub fragment nazwy przystanku: ")
	var searchQuery string
	fmt.Scanln(&searchQuery)

	matchingStops := searchStops(busStops, searchQuery)
	if len(matchingStops) == 0 {
		fmt.Println("Nie znaleziono przystanku o podanej nazwie")
		return BusStop{}
	} 

	for i, stop := range matchingStops {
		fmt.Printf("%d. %s %s\n", i+1, stop.Name, stop.StopCode)
	}

	fmt.Print("Wybierz numer przystanku: ")
	var stopChoice int
	fmt.Scanln(&stopChoice)

	if stopChoice < 1 || stopChoice > len(matchingStops) {
		fmt.Println("Nieprawidłowy wybór")
		return BusStop{}
	}

	chosenStop := matchingStops[stopChoice-1]
	
	return chosenStop
}

func searchStops(busStops []BusStop, query string) []BusStop {
	var results []BusStop
	for _, stop := range busStops {
		fullName := stop.Name + " " + stop.StopCode
		if strings.Contains(strings.ToLower(fullName), strings.ToLower(query)) {
			results = append(results, stop)
		}
	}
	return results
	
}

func formatTimeToLocal(timeStr string) string {
    if timeStr == "" {
        return ""
    }
    
    var parsedTime time.Time
    var err error

    if strings.HasPrefix(timeStr, "1899-12-") {
        parsedTime, err = time.Parse("2006-01-02T15:04:05", timeStr)
    } else {
        parsedTime, err = time.Parse(time.RFC3339, timeStr)
        if err == nil {
            polandLoc, locErr := time.LoadLocation("Europe/Warsaw")
            if locErr == nil {
                parsedTime = parsedTime.In(polandLoc)
            }
        }
    }
    
    if err != nil {
        if tIndex := strings.Index(timeStr, "T"); tIndex >= 0 && tIndex+9 <= len(timeStr) {
            return timeStr[tIndex+1:tIndex+9] 
        }
        
        return timeStr
    }
    
    return parsedTime.Format("15:04:05")
}

func findDeparture(departures []Departure, routeID int, theoreticalTime string) (Departure, error) {
	for _, dep := range departures {
		if dep.RouteID == routeID && formatTimeToLocal(dep.Theoretical) == theoreticalTime {
			return dep, nil
		}
	}
	return Departure{}, fmt.Errorf("Nie znaleziono odjazdu dla trasy %d", routeID)
}

func formatDelayTime(delay int) string {
	minutes := delay / 60
	seconds := delay % 60
	if delay < 0 {
		if minutes > 0 {
			return fmt.Sprintf("Przed czasem: %d m %02d s", minutes, seconds)
		} else {
			return fmt.Sprintf("Przed czasem: %02d s", seconds)
		}
	} else if delay > 0 {
		if minutes > 0 {
			return fmt.Sprintf("Opóźnienie: +%d m %02d s", minutes, seconds)
		} else {
			return fmt.Sprintf("Opóźnienie: +%02d s", seconds)
		}
	}
	return "Punktualnie"
}

func findStop(busStops []BusStop, stopID int) (BusStop, error) {
	for _, stop := range busStops {
		if stop.ID == stopID {
			return stop, nil
		}
	}
	return BusStop{}, fmt.Errorf("Nie znaleziono przystanku o ID %d", stopID)
}

func checkIfEnd(stopTime StopTimes) bool {
	return stopTime.StopSequence == 0 || !stopTime.Passenger
}


func findStopTime(stopTimes []StopTimes, chosenDeparture Departure, chosenStop BusStop, theoricalTime string) int{
	index := -1
	for i, stopTime := range stopTimes {
		departureTimeOnly := formatTimeToLocal(stopTime.DepartureTime) //szukasz przystanku w rozkladzie
		
		if stopTime.RouteId == chosenDeparture.RouteID && 
		stopTime.StopId == chosenStop.ID &&
		departureTimeOnly == theoricalTime {
			
			index = i
			break
		}
	}
	return index
}