package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func main() {
    var wg sync.WaitGroup
    var statsWg sync.WaitGroup  
    
    numberOfOrders := 5
	numberOfWorkers := 2
    processingChanel := make(chan Order)
    outputChannel := make(chan ProcessResult)
    
    for i := 0; i < numberOfWorkers; i++ {
        wg.Add(1)
        go worker(i + 1, &wg, processingChanel, outputChannel)
    }
    
    statsWg.Add(1)
    go func() {
        defer statsWg.Done()
        StatisticsCollector(outputChannel)
    }()
    
	
    for i := 1; i <= numberOfOrders; i++ {
        time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
        order := GenerateRandomOrder()
        processingChanel <- order
    }
    
    close(processingChanel) 
    wg.Wait()                
    close(outputChannel)     
    statsWg.Wait()           
}


func StatisticsCollector(outputChannel chan ProcessResult) {
	processedOrders := []ProcessResult{}
	successCount := 0
	for result := range outputChannel {
		if result.Success {
			successCount++
			fmt.Printf("[Result] Zamówienie ID: %d od %s przetworzono pomyślnie w czasie: %v\n",
				result.OrderID, result.CustomerName, result.ProcessTime)
		} else {
			fmt.Printf("[Result] Błąd przetwarzania zamówienia ID: %d od %s: %v\n",
				result.OrderID, result.CustomerName, result.Error)
		}
		processedOrders = append(processedOrders, result)
	}
	fmt.Printf("[Statistics] %d %% zamówień zakończono sukcesem \n", successCount*100/len(processedOrders))
	fmt.Printf("[Statistics] %d zamówienia zakończono błędem", len(processedOrders)-successCount)

}