use polars::prelude::*;
use polars_core::datatypes::Field;

fn main() {
    let mut schema = Schema::new();
    schema.with_column(
        "Prices".to_owned(),
        DataType::List(Box::from(DataType::Struct(vec![
            Field {
                name: "Country".to_owned(),
                dtype: DataType::Utf8,
            },
            // Field {
            //     name: "Price".to_owned(),
            //     dtype: DataType::UInt32,
            // },
        ]))),
    );

    let res =
        LazyJsonLineReader::new("../../../data-generator/product-data.ndjson".to_owned())
            // .with_batch_size(Some(50))
            .with_schema(schema)
            .finish()
            .unwrap()
            .explode([col("Prices")])
            .groupby(["Prices"])
            .agg([count()])
            .collect();

    println!("{:?}", res);
}
