defmodule Anagram do
  @doc """
  Returns all candidates that are anagrams of, but not equal to, 'base'.
  """
  @spec match(String.t, [String.t]) :: [String.t]
  def match(base, candidates) do
    candidates
    |> Enum.filter(fn(candidate) ->
      String.length(candidate) == String.length(base) end)
    |> Enum.filter(fn(candidate) ->
      Set.subset?(make_set(base), make_set(candidate)) end)
  end

  defp make_set(str) do
    str
    |> String.downcase
    |> String.split("")
    |> Enum.into(MapSet.new)
  end
end
