package main 

type BusStop struct {
	ID int `json:"stopId"`
	Name string `json:"stopName"`
	StopCode string `json:"stopCode"`
	ShortName string `json:"stopShortName"`

}

type BusStopData struct {
    LastUpdate string    `json:"lastUpdate"`
    Stops      []BusStop `json:"stops"` 
}

type Departure struct {
	ID 			string `json:"id"`
	RouteID     int    `json:"routeId"` 
	TripID 		int    `json:"tripId"`     
	Headsign    string `json:"headsign"`            
	Estimated   string `json:"estimatedTime"`       
	Theoretical string `json:"theoreticalTime"`     
	Delay       int    `json:"delayInSeconds"`  
}

type DepartureData struct {
	LastUpdate string      `json:"lastUpdate"`
	Departures []Departure `json:"departures"`
}

type StopTimesData struct {
	LastUpdate string `json:"lastUpdate"`
	StopTimes  []StopTimes `json:"stopTimes"`
}
	

type StopTimes struct {
	RouteId int `json:"routeId"`
	StopId int `json:"stopId"`
	TripId int `json:"tripId"`
	StopSequence int `json:"stopSequence"`
	ArrivalTime string `json:"arrivalTime"`
	DepartureTime string `json:"departureTime"`
	Passenger bool `json:"passenger"`
}