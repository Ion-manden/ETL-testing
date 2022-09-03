package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/transforms/stats"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"

	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/gcs"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
)

type price struct {
	Country string `json:"Country"`
}

type product struct {
	Prices []price `json:"Prices"`
}

func main() {
	flag.Parse()
	// beam.Init() is an initialization hook that must be called on startup.
	beam.Init()

	// Create the Pipeline object and root scope.
	p := beam.NewPipeline()
	s := p.Root()

	lines := textio.Read(s, "../../../data-generator/product-data.ndjson")

	countries := beam.ParDo(s, func(line string, emit func(string)) {
		var p product

		err := json.Unmarshal([]byte(line), &p)
		if err != nil {
			log.Println("Error on Unmarshal :", err)
			return
		}

		for _, price := range p.Prices {
			emit(price.Country)
		}
	}, lines)

	counted := stats.Count(s, countries)

	formatted := beam.ParDo(s, func(w string, c int) string {
		return fmt.Sprintf("%s: %v", w, c)
	}, counted)

	textio.Write(s, "result.txt", formatted)

	if err := beamx.Run(context.Background(), p); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}
