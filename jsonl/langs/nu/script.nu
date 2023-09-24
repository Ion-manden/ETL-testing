open ../../data-generator/product-data.ndjson --raw
| lines
| each { from json }
| each { get Prices }
| flatten
| each { get Country }
| sort
| uniq --count
