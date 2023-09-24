package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/reactivex/rxgo/v2"
)

type price struct {
	Country string `json:"Country"`
}

type product struct {
	Prices []price `json:"Prices"`
}

func main() {
	file, err := os.Open("../../../data-generator/product-data.ndjson")
	if err != nil {
		log.Fatal("Error on Open :", err)
	}

	ch := make(chan rxgo.Item)
	reader := bufio.NewReader(file)
	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Println("Error on ReadLine :", err)
				continue
			}

			item := rxgo.Item{V: line}
			ch <- item
		}

		close(ch)
	}()

	observable := rxgo.FromChannel(ch)
	result, err := observable.
		Map(func(_ context.Context, item interface{}) (interface{}, error) {
			var p product

			err = json.Unmarshal([]byte(item.(string)), &p)
			if err != nil {
				log.Println("Error on Unmarshal :", err)
				return nil, err
			}

			return p, nil
		}, rxgo.WithPool(32)).
		FlatMap(func(i rxgo.Item) rxgo.Observable {
			return rxgo.Just(i.V.(product).Prices)()
		}, rxgo.WithPool(32)).
		Map(func(_ context.Context, item interface{}) (interface{}, error) {
			return item.(price).Country, nil
		}, rxgo.WithPool(32)).
		Reduce(func(_ context.Context, acc interface{}, elem interface{}) (interface{}, error) {
			if acc == nil {
				return map[string]int{}, nil
			}

			country := elem.(string)

			acc.(map[string]int)[country] = acc.(map[string]int)[country] + 1

			return acc, nil
		}).
		Get()

	if err != nil {
		log.Fatal("Error on observable :", err)
	}

	b, _ := json.Marshal(result.V)
	fmt.Println(string(b))
}
