defmodule DayFour do

  def main(input_file) do
    IO.puts("AoC Day 4 2021")
    {:ok, board} = parse_bingo_input(input_file)

    IO.puts("parsed bingo board: #{inspect(board)}")
  end

  defp parse_bingo_input(input_file) do
    case File.read(input_file) do
      {:ok, contents} ->
        [instructions | unstructured_boards] = String.split(contents, "\n")
        {:ok, %{
          instructions: instructions,
          boards: unstructured_boards
                  |> Enum.chunk_by(&(&1 == ""))
                  |> Enum.filter(&(&1 != [""]))
                  |> Enum.map(fn board ->
                    board
                    |> Enum.with_index()
                    |> Enum.map(fn {row, row_number} ->
                      {row_number,
                       row
                       |> String.split(" ", trim: true)
                       |> Enum.with_index()
                       |> Enum.map(fn {v,k} -> {k,String.to_integer(v)} end)
                       |> Enum.into(%{})}
                    end)
                    |> Enum.into(%{})
                  end)
          }
        }
      err -> err
    end
  end

end

DayFour.main("input/input4.txt")
