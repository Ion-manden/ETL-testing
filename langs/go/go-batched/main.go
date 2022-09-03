package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type price struct {
	Country string `json:"Country"`
}

type product struct {
	Prices []price `json:"Prices"`
}

func main() {
	workChan := make(chan []string, 10)
	resultsChan := make(chan map[string]int, 5)

	workerWg := sync.WaitGroup{}
	workerCount := 5
	for i := 0; i < workerCount; i++ {
		workerWg.Add(1)
		go func() {
			startBatchHandler(workChan, resultsChan)
			workerWg.Done()
		}()
	}

	file, err := os.Open("../../../data-generator/product-data.ndjson")
	if err != nil {
		log.Fatal("Error on Open :", err)
	}

	reader := bufio.NewReader(file)

	countryCountsResult := map[string]int{}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for partialCountryCount := range resultsChan {
			for country, count := range partialCountryCount {
				currentCount, ok := countryCountsResult[country]
				if !ok {
					currentCount = 0
				}

				countryCountsResult[country] = currentCount + count
			}
		}

		wg.Done()
	}()

	workBuffer := []string{}
	workBufferSize := 100
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				workChan <- workBuffer

				workBuffer = []string{}

				break
			}
			log.Println("Error on ReadLine :", err)
			continue
		}

		workBuffer = append(workBuffer, line)

		if len(workBuffer) >= workBufferSize {
			workChan <- workBuffer

			workBuffer = []string{}
		}
	}

	close(workChan)
	workerWg.Wait()
	close(resultsChan)
	wg.Wait()

	b, _ := json.Marshal(countryCountsResult)
	fmt.Println(string(b))
}

func startBatchHandler(in <-chan []string, out chan<- map[string]int) {
	for buffer := range in {
		countryCounts := map[string]int{}

		for _, line := range buffer {
			var p product

			err := json.Unmarshal([]byte(line), &p)
			if err != nil {
				log.Println("Error on Unmarshal :", err)
				continue
			}

			for _, price := range p.Prices {
				countryCounts[price.Country] = countryCounts[price.Country] + 1
			}
		}

		out <- countryCounts
	}
}
