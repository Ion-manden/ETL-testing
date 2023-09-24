use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;

use rayon::prelude::*;

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

    let country_count: HashMap<String, u32> = reader
        .lines()
        .filter_map(|line: Result<String, _>| line.ok())
        .par_bridge() // Changes from iterator to parallel iterator
        .filter_map(|line: String| serde_json::from_str(&line).ok()) // filter out bad lines
        .flat_map(|product: Product| product.Prices)
        .map(|price: Price| price.Country)
        .fold(
            || HashMap::new(),
            |mut country_count: HashMap<String, u32>, country: String| {
                *country_count.entry(country).or_insert(0) += 1;

                country_count
            },
        )
        .reduce(
            || HashMap::new(),
            |mut result, word_count| {
                merge_maps(&mut result, word_count);
                result
            },
        );

    println!("{:?}", country_count);

    Ok(())
}

fn merge_maps(to: &mut HashMap<String, u32>, from: HashMap<String, u32>) {
    for (key, val) in from.iter() {
        let count = to.entry(key.to_owned()).or_insert(0);
        *count += *val;
    }
}
