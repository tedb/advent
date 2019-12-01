defmodule Day1 do
  def fuel(m) when div(m, 3)-2 <= 0 do
    0
  end

  def fuel(m) do
    addl = div(m, 3)-2
    addl + fuel(addl)
  end

  def a do
    File.stream!("input")
    |> Enum.map(&String.to_integer(String.trim(&1)))
    |> Enum.map(&( div(&1, 3) - 2))
    |> Enum.sum
  end

  def b do
    File.stream!("input")
    |> Enum.map(&String.to_integer(String.trim(&1)))
    |> Enum.map(&fuel(&1))
    |> Enum.sum
  end
end

IO.puts "day1a: #{Day1.a}\nday1b: #{Day1.b}"
