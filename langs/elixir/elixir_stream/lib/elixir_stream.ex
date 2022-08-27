defmodule ElixirStream do
  use Application

  def start(_type, _args) do
    stream()
    {:ok, self()}
  end

  def stream() do
    File.stream!("../../../data-generator/product-data.ndjson")
    |> Stream.flat_map(fn l ->
      {:ok, product} = Jason.decode(l)
      product["Prices"]
    end)
    |> Stream.map(fn p ->
      p["Country"]
    end)
    |> Enum.reduce(%{}, fn c, acc ->
      Map.update(
        acc,
        c,
        1,
        fn c -> c + 1 end
      )
    end)
    |> Jason.encode!()
    |> IO.puts()
  end
end
