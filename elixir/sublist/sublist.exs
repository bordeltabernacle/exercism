defmodule Sublist do

  @doc """
  Returns whether the first list is a sublist
  or a superlist of the second list and if not
  whether it is equal or unequal to the second list.
  """
  @spec compare(list(), list()) :: :atom
  def compare(list_a, list_b) do
    {len_a, len_b} = {length(list_a), length(list_b)}
    cond do
      list_a == list_b ->
        :equal
      len_a < len_b && sublist?(list_a, list_b, len_a) ->
        :sublist
      list_a > len_b && sublist?(list_b, list_a, len_b) ->
        :superlist
      list_a != list_b ->
        :unequal
    end
  end

  defp sublist?([], _, _), do: :true
  defp sublist?(_, [], _), do: :false
  defp sublist?(smaller, larger = [_|tail_b], len_smaller) do
    cond do
      length(tail_b) + 1 < len_smaller ->
        :false
      smaller === larger |> Enum.take(len_smaller) |> Enum.to_list ->
        :true
      true ->
        sublist?(smaller, tail_b, len_smaller)
    end
  end

end
