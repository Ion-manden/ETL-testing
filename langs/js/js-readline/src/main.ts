import { createReadStream } from 'fs';
import readline from 'readline';

type Product = {
  Id: number;
  Name: string;
  Description: string;
  Prices: Price[];
  Created: Date;
  CreatedFormat: Date;
}

type Price = {
  Country: string;
  Price: number;
}


const countryCount: Record<string, number> = {};

readline.createInterface({
  input: createReadStream('../../../data-generator/product-data.ndjson'),
  terminal: false
}).on('line', function(line) {
  const product: Product = JSON.parse(line);
  const prices = product.Prices;

  for (const price of prices) {
    const country = price.Country;

    countryCount[country] = (countryCount[country] ?? 0) + 1;
  }
}).on('close', () => {
  console.log(countryCount)
});

