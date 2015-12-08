defmodule Sublist do
  @doc """
  Returns whether the first list is a sublist or a superlist of the second list
  and if not whether it is equal or unequal to the second list.
  """
  @spec compare(list, list) :: atom
  def compare(a, b) do
    compare(a, b, length(a), length(b))
  end

  defp compare(a, b, la, lb) do
    cond do
      a == b && la == lb ->
        :equal
      la < lb && sublist?(a, b, la, lb) ->
        :sublist
      la > lb && sublist?(b, a, la, lb) ->
        :superlist
      a != b ->
        :unequal
    end
  end

  def sublist?([], _, _, _), do: :true
  def sublist?(_, [], _, _), do: :false
  def sublist?(a = [ha|ta], b = [hb|tb], la, lb) do
    cond do
      a == b ->
        :true
      ha === hb ->
        sublist?(ta, tb, la, lb)
      ha !== hb ->
        sublist?(a, tb, la, lb)
    end
  end

end
