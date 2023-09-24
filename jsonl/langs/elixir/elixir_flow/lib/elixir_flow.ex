defmodule ElixirFlow do
  use Application

  def start(_type, _args) do
    flow()
    {:ok, self()}
  end

  def flow() do
    File.stream!("../../../data-generator/product-data.ndjson")
    |> Flow.from_enumerable()
    |> Flow.flat_map(fn l ->
      {:ok, product} = Jason.decode(l)
      product["Prices"]
    end)
    |> Flow.map(fn p ->
      p["Country"]
    end)
    |> Flow.partition()
    |> Flow.reduce(
      fn -> %{} end,
      fn c, acc ->
        Map.update(
          acc,
          c,
          1,
          fn c -> c + 1 end
        )
      end
    )
    |> Enum.into(%{})
    |> Jason.encode!()
    |> IO.puts()
  end
end
