defmodule Sublist do

  @doc """
  Returns whether the first list is a sublist
  or a superlist of the second list and if not
  whether it is equal or unequal to the second list.
  """
  @spec compare(list(), list()) :: :atom
  def compare(a, b) do
    {la, lb} = {length(a), length(b)}
    cond do
      a == b ->
        :equal
      la < lb && sublist?(a, b, la) ->
        :sublist
      la > lb && sublist?(b, a, lb) ->
        :superlist
      a != b ->
        :unequal
    end
  end

  defp sublist?([], _, _), do: :true
  defp sublist?(_, [], _), do: :false
  defp sublist?(smaller, larger = [_|tb], len_smaller) do
    cond do
      smaller === Enum.take(larger, len_smaller) ->
        :true
      true ->
        sublist?(smaller, tb, len_smaller)
    end
  end

end
