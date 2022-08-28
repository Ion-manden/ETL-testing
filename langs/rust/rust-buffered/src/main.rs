use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;

use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct Product {
    Prices: Vec<Price>,
}

#[derive(Serialize, Deserialize)]
struct Price {
    Country: String,
}

fn main() -> std::io::Result<()> {
    let f = File::open("../../../data-generator/product-data.ndjson")?;
    let reader = BufReader::new(f);

    let mut country_count: HashMap<String, i32> = HashMap::new();

    for line in reader.lines() {
        if let Ok(json) = line {
            let data: &str = &json[..];
            // Parse the string of data into serde_json::Value.
            // let f: Value = serde_json::from_str(data)?;
            let product: Product = serde_json::from_str(data)?;

            for price in product.Prices {
                let prev_count = country_count.get(&price.Country).unwrap_or(&0).to_owned();

                country_count.insert(price.Country, prev_count + 1);
            }
        }
    }

    println!("{:?}", country_count);

    Ok(())
}
