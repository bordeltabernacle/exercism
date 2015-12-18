defmodule Anagram do
  @doc """
  Returns all candidates that are anagrams of, but not equal to, 'base'.
  """
  @spec match(String.t, [String.t]) :: [String.t]
  def match(base, candidates) do
    Enum.filter candidates,
    &(sorted(String.downcase(&1)) == sorted(String.downcase(base))
      && String.downcase(&1) != String.downcase(base))
  end

  defp sorted(str) do
    str
    |> String.graphemes
    |> Enum.sort
  end
end
