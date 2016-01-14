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

  defp validate(number) do
    number_length = number |> to_char_list |> length
    cond do
      number_length == 10 ->
        number
      number_length < 10 ->
        @bad_number
      number_length == 11 && String.starts_with?(number, "1") ->
        String.slice(number, 1..11)
      number_length > 10 ->
        @bad_number
    end
  end

  @doc """
  Extract the area code from a phone number
  """
  @spec area_code(String.t) :: String.t
  def area_code(raw) do
    raw
    |> number
    |> String.slice(0..2)
  end

  @doc """
  Extract the prefix code from a phone number
  """
  @spec prefix(String.t) :: String.t
  def prefix(raw) do
    raw
    |> number
    |> String.slice(3..5)
  end

  @doc """
  Extract the line number from a phone number
  """
  @spec line(String.t) :: String.t
  def line(raw) do
    raw
    |> number
    |> String.slice(6..9)
  end

  @doc """
  Pretty print a phone number
  """
  @spec pretty(String.t) :: String.t
  def pretty(raw) do
    "(#{area_code(raw)}) #{prefix(raw)}-#{line(raw)}"
  end
end
