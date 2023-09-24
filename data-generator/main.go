package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Product struct {
	Id            int
	Name          string  `fake:"{appname}"`
	Description   string  `fake:"{paragraph:5,12,8}"`
	Prices        []Price `fakesize:"12"`
	Created       time.Time
	CreatedFormat time.Time `fake:"{year}-{month}-{day}" format:"2006-01-02"`
}

type Price struct {
	Country string  `fake:"{country}"`
	Price   float64 `fake:"{price:1,50000}"`
}

func main() {
	fileLengthArg := os.Args[1]
	fileLength, _ := strconv.Atoi(fileLengthArg)
	f, _ := os.Create("product-data.ndjson")
	defer f.Close()

	writer := bufio.NewWriter(f)

	for fl := 0; fl < fileLength; fl++ {
		var p Product
		gofakeit.Struct(&p)

		b, _ := json.Marshal(p)

		_, err := writer.WriteString(fmt.Sprint(string(b), "\n"))
		if err != nil {
			log.Println(err)
		}
	}

	writer.Flush()
}
