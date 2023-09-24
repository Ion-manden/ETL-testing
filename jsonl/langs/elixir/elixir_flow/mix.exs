defmodule ElixirFlow.MixProject do
  use Mix.Project

  def project do
    [
      app: :elixir_flow,
      version: "0.1.0",
      elixir: "~> 1.13",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      mod: {ElixirFlow, []},
      extra_applications: [:logger]
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:flow, "~> 1.2"},
      {:jason, "~> 1.2"}
    ]
  end
end
