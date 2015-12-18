defmodule Anagram do
  @doc """
  Returns all candidates that are anagrams of, but not equal to, 'base'.
  """
  @spec match(String.t, [String.t]) :: [String.t]
  def match(base, candidates) do
    candidates
    |> Enum.filter(&(sorted(&1) == sorted(base) && different?(&1, base)))
  end

  defp sorted(str) do
    str
    |> String.downcase
    |> String.graphemes
    |> Enum.sort
  end

  defp different?(str_a, str_b) do
    String.downcase(str_a) != String.downcase(str_b)
  end
end
