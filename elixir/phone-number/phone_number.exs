defmodule Phone do

  @valid_number_regex ~r"^\D*?1?\D*?(\d{3})\D*(\d{3})\D*(\d{4})$"

  @doc """
  Remove formatting from a phone number.
  Returns "0000000000" if phone number is not valid
  (10 digits or "1" followed by 10 digits)
  """
  @spec number(String.t) :: String.t
  def number(raw) do
    raw
    |> to_number_list
    |> List.to_string
  end

  defp to_number_list(raw) do
    @valid_number_regex
    |> Regex.run(raw, [capture: :all_but_first]) || ["000", "000", "0000"]
  end

  @doc """
  Extract the area code from a phone number
  Returns the first three digits from a phone number,
  ignoring long distance indicator
  """
  @spec area_code(String.t) :: String.t
  def area_code(raw) do
    raw
    |> to_number_list
    |> List.first
  end

  @doc """
  Pretty print a phone number
  Wraps the area code in parentheses and separates
  exchange and subscriber number with a dash.
  """
  @spec pretty(String.t) :: String.t
  def pretty(raw) do
    [area_code, prefix, main] = to_number_list(raw)
    "(#{area_code}) #{prefix}-#{main}"
  end
end
