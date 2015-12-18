defmodule Anagram do
  @doc """
  Returns all candidates that are anagrams of, but not equal to, 'base'.
  """
  @spec match(String.t, [String.t]) :: [String.t]
  def match(base, candidates) do
    candidates
    |> Enum.filter(fn(candidate) ->
      sorted(String.downcase(candidate)) == sorted(String.downcase(base))
      && String.downcase(candidate) != String.downcase(base)
    end)
  end

  defp sorted(str) do
    str
    |> String.graphemes
    |> Enum.sort
  end
end
