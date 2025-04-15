package main 

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup, processingChanel chan Order, outputChannel chan ProcessResult) {
	defer wg.Done()
	for order := range processingChanel {
		
		fmt.Printf("[Worker %d] Przetwarzam zamówienie: %+v\n", id, order) 
		result := OrderProcessing(order)
		//for zamienic na if i wtedy statystyka bedzie miala sens
		for !result.Success {
			fmt.Printf("[Worker %d] Błąd przetwarzania zamówienia ponawiam probe: %+v\n", id, result)
			result = OrderProcessing(order)
		}
		outputChannel <- result
		fmt.Printf("[Worker %d] Przetworzone zamówienie: %+v\n", id, result)
	}
}