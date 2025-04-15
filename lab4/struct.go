package main

import (
	"time"
	"math/rand"
	"fmt"
)


type Order struct {
	ID           int
	CustomerName string
	Items        []string
	TotalAmount  float64
}


type ProcessResult struct {
	OrderID      int
	CustomerName string
	Success      bool
	ProcessTime  time.Duration
	Error        error
}

func GenerateRandomOrder() Order {
	customers := []string{"Krystyna", "Paweł", "Lexy", "Wojciech", "Igor"}
	items := []string{"Pędzel", "Korniszony", "Krowa", "Banan", "Błyszczyk", "Zaliczenie z golanga"}
	prices := []float64{9.99, 6.49, 200.13, 1.23, 39.99, 500}


	id := rand.Intn(100) +1
	customer := customers[rand.Intn(len(customers))]
	itemAmount := rand.Intn(5) + 1
	var ordered []string
	var total float64
	for i:= 0; i < itemAmount; i++ {
		index := rand.Intn(len(items))
		ordered = append(ordered, items[index])
		total += prices[index]
	}

	return Order {
		ID: id,
		CustomerName: customer,
		Items: ordered,
		TotalAmount: total,
	}
}

func OrderProcessing(order Order) ProcessResult { 
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	success := rand.Float32() < 0.6
	var err error
	if !success {
		err = fmt.Errorf("error processing order %d", order.ID)
	}
	return ProcessResult{
		OrderID:      order.ID,
		CustomerName: order.CustomerName,
		Success:      success,
		ProcessTime:  time.Since(start),
		Error:        err,
	}

}