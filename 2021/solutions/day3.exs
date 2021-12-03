defmodule DayThree do

  def main(input_file) do
    diagnostic_report = parse_input(input_file)
    IO.puts("Day Three AoC")

    analysis = analyze_report(diagnostic_report)

    IO.puts("Part one: #{get_gamma_rate(analysis) * get_epsilon_rate(analysis)}")

    IO.puts("Part two: #{get_oxygen_generator_rating(diagnostic_report) * get_co2_scrubber_rating(diagnostic_report)}")
  end

  def parse_input(input_file) do
    File.read!(input_file)
    |> String.split("\n", trim: true)
  end

  defp analyze_report(report) do
    Enum.reduce(report, %{}, fn line, acc ->
      line
      |> String.to_charlist()
      |> Enum.with_index()
      |> Enum.reduce(acc, fn {bit, index}, acc ->
        Map.update(acc, index, %{bit => 1}, fn existing_record ->
          existing_record |> Map.update(bit, 1, fn current_count -> current_count + 1 end)
        end)
      end)
    end)
  end

  defp get_gamma_rate(analysis) do
    get_simple_rate(analysis, &Enum.max_by/2)
  end

  defp get_epsilon_rate(analysis) do
    get_simple_rate(analysis, &Enum.min_by/2)
  end

  defp get_simple_rate(analysis, filter_function) do
    analysis
    |> Map.keys()
    |> Enum.sort()
    |> Enum.reduce([], fn index, acc ->
      [analysis[index]
       |> Map.to_list()
       |> (&(filter_function.(&1, fn val -> val |> elem(1) end))).()
       |> elem(0) | acc]
    end)
    |> Enum.reverse()
    |> to_string()
    |> String.to_integer(2)
  end

  def get_oxygen_generator_rating(report) do
    get_complex_rating(report, &Enum.max_by/2, ?1, "", 0)
  end

  def get_co2_scrubber_rating(report) do
    get_complex_rating(report, &Enum.min_by/2, ?0, "", 0)
  end

  def get_complex_rating(report, filter_function, default_value, accumulator, position) do
    analysis = analyze_report(report)
    next_value = if (analysis[position] |> Map.to_list() |> Enum.uniq_by(fn {_key, val} -> val end) |> length()) > 1 do
      analysis[position] |> Map.to_list() |> (&(filter_function.(&1, fn val -> val |> elem(1) end))).() |> elem(0)
    else
      default_value
    end
    next_accumulator = accumulator <> <<next_value>>

    case length(report) do
      1 ->
        [final_value] = report
        final_value |> String.to_integer(2)
      _ ->
        get_complex_rating(
          report |> Enum.filter(fn val -> String.starts_with?(val, next_accumulator) end),
          filter_function,
          default_value,
          next_accumulator,
          position + 1
        )
    end
  end
end

DayThree.main("input/input3.txt")
