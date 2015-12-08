defmodule Sublist do
  @doc """
  Returns whether the first list is a sublist or a superlist of the second list
  and if not whether it is equal or unequal to the second list.
  """
  @spec compare(list, list) :: atom
  def compare(a, b) do
    {la, lb} = {length(a), length(b)}
    cond do
      a == b ->
        :equal
      la < lb && sublist?(a, b, a) ->
        :sublist
      la > lb && sublist?(b, a, b) ->
        :superlist
      a != b ->
        :unequal
    end
  end

  def sublist?([], _, _), do: :true
  def sublist?(_, [], _), do: :false
  def sublist?(a = [ha|ta], b = [hb|tb], init) do
    cond do
      a == b ->
        :true
      ha === hb ->
        sublist?(ta, tb, init)
      ha !== hb ->
        sublist?(init, tb, init)
    end
  end

end
