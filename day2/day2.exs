defmodule Day2 do
  def a do
    File.read!("input")
    |> String.trim()
    |> String.split(",")
    |> Enum.map(&String.to_integer(&1))
    |> List.replace_at(1, 12)
    |> List.replace_at(2, 2)
    |> op(0)
    |> IO.puts
  end

  def b do
    File.read!("input")
    |> String.trim()
    |> String.split(",")
    |> Enum.map(&String.to_integer(&1))
    |> op_seek
    |> IO.puts
  end

  def op_seek(list) do
    # todo: list comprehension: loop through 1..100 twice, op_with(a, b), filter for result == 19690720,
    # return (a * 100) + b
    a = 12
    b = 2
    op_with(list, a, b)
    a * 100 + b
  end

  def op_with([op, _, _ | rest], a, b) do
    op([op, a, b | rest], 0)
  end

  def op(prog) do
    op(prog, 0)
  end

  def op(prog, pc) do
    IO.inspect prog
    case Enum.at(prog, pc) do
      1 ->
        v = Enum.at(prog, Enum.at(prog, pc+1)) + Enum.at(prog, Enum.at(prog, pc+2))
        op(List.replace_at(prog, Enum.at(prog, pc+3), v), pc+4)
      2 ->
        v = Enum.at(prog, Enum.at(prog, pc+1)) * Enum.at(prog, Enum.at(prog, pc+2))
        op(List.replace_at(prog, Enum.at(prog, pc+3), v), pc+4)
      99 ->
        Enum.at(prog, 0)
    end
  end
end
