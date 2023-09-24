defmodule ElixirStreamTest do
  use ExUnit.Case
  doctest ElixirStream

  test "greets the world" do
    assert ElixirStream.hello() == :world
  end
end
