package main

import (
	"fmt"
	"time"

	"math/rand"

)


func GenerujPESEL(birthDate time.Time, gender string) [11]int {

	var cyfryPESEL [11]int 

	year := birthDate.Year()
	month := int(birthDate.Month())
	day := birthDate.Day()

	cyfryPESEL[0] = (year / 10)%10
	cyfryPESEL[1] = year % 10

	switch {
		case 1800 <= year && year <= 1899:
			month += 80
		case 2000 <= year && year <= 2099:
			month += 20
		case 2100 <= year && year <= 2199:
			month += 40
		case 2200 <= year && year <= 2299:
			month += 60
	}
	cyfryPESEL[2] = month / 10
	cyfryPESEL[3] = month % 10

	cyfryPESEL[4] = day / 10
	cyfryPESEL[5] = day % 10

	randomSerial := rand.Intn(900) + 100 

	cyfryPESEL[6] = randomSerial / 100
	cyfryPESEL[7] = (randomSerial / 10) % 10
	cyfryPESEL[8] = randomSerial % 10

	switch gender{
	case "M":
		cyfryPESEL[9] = rand.Intn(5) * 2 + 1
	case "K":
		cyfryPESEL[9] = rand.Intn(5) * 2 
	}

	var wagi = [4]int{1, 3, 7, 9}

	var suma int = 0
	for i := 0; i < 10; i++ { //tu rangem
		liczba := cyfryPESEL[i] * wagi[i%4] % 10
		suma += liczba
	}

	suma = suma % 10
	cyfryPESEL[10] = 10 -suma

	return cyfryPESEL
}


func WeryfikujPESEL(cyfryPESEL [11]int) bool {

	var czyPESEL bool
	var wagi = [4]int{1, 3, 7, 9}
	var suma int = 0
	for i := 0; i < 10; i++ {
		liczba := cyfryPESEL[i] * wagi[i%4] % 10
		suma += liczba
	}

	czyPESEL = (10-(suma%10) == cyfryPESEL[10])

	return czyPESEL
}

func main() { 
	rand.Seed(time.Now().UnixNano())

	birthDate := time.Date(1980, 2, 26, 0, 0, 0, 0, time.UTC)

	pesel := GenerujPESEL(birthDate, "M")
	
	fmt.Println("Wygenerowany PESEL:", pesel)

	fmt.Println("Czy numer PESEL jest poprawny:", WeryfikujPESEL(pesel))


}

