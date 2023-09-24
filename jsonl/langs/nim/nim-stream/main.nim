import std/json
from tables import initCountTable, inc
import std/streams

var strm = newFileStream("../../../data-generator/product-data.ndjson", fmRead)

var line = ""

var countryCounts = initCountTable[string]()

if not isNil(strm):
  while strm.readLine(line):
    let product = parseJson(line)

    for price in product["Prices"]:
      var country = price["Country"].getStr()
      countryCounts.inc country

  strm.close()


