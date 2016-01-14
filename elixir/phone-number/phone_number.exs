defmodule Phone do

  @bad_number "0000000000"

  @doc """
  Remove formatting from a phone number.
  Returns "0000000000" if phone number is not valid
  (10 digits or "1" followed by 10 digits)
  """
  @spec number(String.t) :: String.t
  def number(raw) do
    raw
    |> scrub_number
    |> validate
  end

  defp scrub_number(raw), do: Regex.replace(~r"\W", raw, "")

  defp split_number(number) do
    [String.slice(number, 0..2),
     String.slice(number, 3..5),
     String.slice(number, 6..9)]
  end

  defp validate(number) when byte_size(number) == 10, do: number
  defp validate("1" <> last_ten = number) when byte_size(number) == 11, do: last_ten
  defp validate(_), do: @bad_number

  @doc """
  Extract the area code from a phone number
  """
  @spec area_code(String.t) :: String.t
  def area_code(raw) do
    raw
    |> number
    |> split_number
    |> hd
  end

  @doc """
  Pretty print a phone number
  """
  @spec pretty(String.t) :: String.t
  def pretty(raw) do
    [area_code, prefix, line] = raw |> number |> split_number
    "(#{area_code}) #{prefix}-#{line}"
  end
end
