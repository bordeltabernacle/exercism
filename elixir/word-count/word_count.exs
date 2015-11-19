defmodule Words do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t) :: map()

  def count(sentence) do
    sentence
    |> String.downcase
    |> remove_punctuation
    |> String.split
    |> count(%{})
  end

  def count([], map) do
    map
  end

  def count([head|tail], map) do
    count(tail, Map.update(map, head, 1, &(&1 + 1)))
  end

  defp remove_punctuation(word) do
    String.replace(word, ~r/[^\w\s-]+|_/u, " ")
  end

end
