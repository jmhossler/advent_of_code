defmodule DayTwo do

  def main(input_file) do
    IO.puts("Day Two (AOC)")
    {:ok, instructions} = get_instructions(input_file)

    %{horizontal_position: ending_horizontal_position, depth: ending_depth} =
      Enum.reduce(instructions, %{horizontal_position: 0, depth: 0}, &handle_instruction/2)

    IO.puts("Part one: #{ending_horizontal_position * ending_depth}")

    %{horizontal_position: ending_horizontal_position, depth: ending_depth} =
      Enum.reduce(instructions, %{horizontal_position: 0, depth: 0, aim: 0}, &handle_instruction/2)

    IO.puts("Part two: #{ending_horizontal_position * ending_depth}")
  end

  defp get_instructions(input_file) do
    case File.read(input_file) do
      {:ok, contents} when is_binary(contents) ->
        {:ok,
         contents
         |> String.split("\n", trim: true)
         |> Enum.map(&String.split(&1, " ", trim: true))
         |> Enum.map(fn [direction, string_magnitude] ->
           [direction, String.to_integer(string_magnitude)]
         end)}
      error -> error
    end
  end

  defp handle_instruction(["forward", magnitude], %{horizontal_position: x, depth: y, aim: z}) do
    %{horizontal_position: x + magnitude, depth: y + magnitude * z, aim: z}
  end

  defp handle_instruction(["forward", magnitude], %{horizontal_position: x, depth: y}) do
    %{horizontal_position: x + magnitude, depth: y}
  end

  defp handle_instruction(["up", magnitude], %{horizontal_position: x, depth: y, aim: z}) do
    %{horizontal_position: x, depth: y, aim: z - magnitude}
  end

  defp handle_instruction(["up", magnitude], %{horizontal_position: x, depth: y}) do
    %{horizontal_position: x, depth: y - magnitude}
  end

  defp handle_instruction(["down", magnitude], %{horizontal_position: x, depth: y, aim: z}) do
    %{horizontal_position: x, depth: y, aim: z + magnitude}
  end

  defp handle_instruction(["down", magnitude], %{horizontal_position: x, depth: y}) do
    %{horizontal_position: x, depth: y + magnitude}
  end
end

DayTwo.main("input/input2.txt")
