package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type price struct {
	Country string `json:"Country"`
}

type product struct {
	Prices []price `json:"Prices"`
}

func main() {
	file, err := os.Open("../../../data-generator/short-product-data.ndjson")
	if err != nil {
		log.Fatal("Error on Open :", err)
	}

	reader := bufio.NewReader(file)

	countryCounts := map[string]int{}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error on ReadLine :", err)
			continue
		}

		var p product

		err = json.Unmarshal([]byte(line), &p)
		if err != nil {
			log.Println("Error on Unmarshal :", err)
			continue
		}

		for _, price := range p.Prices {

			countryCounts[price.Country] = countryCounts[price.Country] + 1
		}
	}

	b, _ := json.Marshal(countryCounts)
	fmt.Println(string(b))
}
