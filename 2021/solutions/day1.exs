defmodule DayOne do 
  def main(input_file) do
    {:ok, sea_depths} = get_sea_depths(input_file)

    %{increases: part_one_answer} = Enum.reduce(sea_depths, %{last: nil, increases: 0}, fn e, acc ->
      if !is_nil(acc.last) and acc.last < e do
        %{last: e, increases: acc.increases + 1}
      else
        %{last: e, increases: acc.increases}
      end
    end)

    IO.puts("Part one: #{part_one_answer}")

    %{increases: part_two_answer} = Enum.chunk_every(sea_depths, 3, 1, :discard)
    |> Enum.map(&Enum.sum(&1))
    |> Enum.reduce(%{last: nil, increases: 0}, fn e, acc ->
      if !is_nil(acc.last) and acc.last < e do
        %{last: e, increases: acc.increases + 1}
      else
        %{last: e, increases: acc.increases}
      end
    end)

    IO.puts("Part two: #{part_two_answer}")
  end

  defp get_sea_depths(file_path) do
    case File.read(file_path) do
      {:ok, contents} when is_binary(contents) ->
        {:ok, contents |> String.split("\n", trim: true) |> Enum.map(&String.to_integer(&1))}
      error -> error
    end
  end
end

DayOne.main("input/input1.txt")
