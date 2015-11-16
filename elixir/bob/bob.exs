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
      String.strip(input) == "" ->
				"Fine. Be that way!"
			String.ends_with?(input, "?") ->
				"Sure."
			input == String.upcase(input) and
			Regex.match?(~r/\p{L}+/, input)  ->
				"Whoa, chill out!"
      true ->
				"Whatever."
		end
  end

end
