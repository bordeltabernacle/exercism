defmodule Teenager do

  @moduledoc """
  Bob is a lackadaisical teenager.
  In conversation, his responses are very limited.

  Bob answers 'Sure.' if you ask him a question.
  He answers 'Whoa, chill out!' if you yell at him.
  He says 'Fine. Be that way!' if you address him,
  without actually saying anything.
  He answers 'Whatever.' to anything else.
  """

  @doc """
  Given a sentence as input, returns the appropriate response of Bob,
  the lackadaisical teenager.
  """

  @spec hey(char_list()) :: char_list()

  def hey(input) do
    cond do
      silence?(input)             -> "Fine. Be that way!"
      question?(input)            -> "Sure."
      shouting?(input) and has_alpha?(input) -> "Whoa, chill out!"
      true                        -> "Whatever."
    end
  end

  def silence?(input) do
    String.strip(input) == ""
  end

  def question?(input) do
    String.ends_with?(input, "?")
  end

  def shouting?(input) do
    input == String.upcase(input) and has_alpha?(input)
  end

  def has_alpha?(input) do
    Regex.match?(~r/\p{L}+/, input)
  end

end
