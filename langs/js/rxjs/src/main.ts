import { filter, from, map, mergeMap, reduce, fromEventPattern } from 'rxjs';
import { createReadStream } from 'fs'
import readline from 'readline'

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

async function main(): Promise<void> {
  console.time('exec-time');

  const rl = readline.createInterface({
    input: createReadStream('../../../data-generator/short-product-data.ndjson'),
    terminal: true
  });

  await new Promise<void>((resolve) => {
    fromEventPattern((handler) => rl.on('line', handler), (handler) => rl.on('end', handler))
      .pipe(
        filter((chunk): chunk is string => typeof chunk === 'string'),
        map((line): Product => JSON.parse(line)),
        mergeMap((product) => from(product.Prices)),
        map((price) => price.Country),
        reduce<string, Record<string, number>>(
          (countryCount, country) => {
            countryCount[country] = (countryCount[country] ?? 0) + 1
            return countryCount
          },
          {}
        )
      )
      .subscribe((result) => {
        console.log(result)
      })
      .add(resolve);
  });

  console.timeEnd('exec-time');
}

main().catch(console.error);

