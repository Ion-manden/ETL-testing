const { createReadStream } = require('fs');
var readline = require('readline');

const countryCount = {};

readline.createInterface({
  input: createReadStream('../../../data-generator/product-data.ndjson'),
  terminal: false
}).on('line', function(line) {
  const { Prices } = JSON.parse(line);

  for (const price of Prices) {
    const country = price.Country;

    countryCount[country] = (countryCount[country] ?? 0) + 1;
  }
}).on('close', () => {
  console.log(countryCount)
});

