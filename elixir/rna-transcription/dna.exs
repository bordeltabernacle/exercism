defmodule DNA do

  @dna_to_rna %{71 => 'C', 67 => 'G', 84 => 'A', 65 => 'U'}

  @doc """
  Transcribes a character list representing DNA nucleotides to RNA

  ## Examples

  iex> DNA.to_rna('ACTG')
  'UGAC'
  """
  @spec to_rna([char]) :: [char]
  def to_rna(dna) do
    dna
    |> Enum.reduce('', &(&2 ++ converter(&1)))
  end

  defp converter(dna) do
    %{^dna => result} = @dna_to_rna
    result
  end
end
