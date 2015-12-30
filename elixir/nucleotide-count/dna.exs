defmodule DNA do
  @nucleotides [?A, ?C, ?G, ?T]

  @doc """
  Counts individual nucleotides in a DNA strand.

  ## Examples

  iex> DNA.count('AATAA', ?A)
  4

  iex> DNA.count('AATAA', ?T)
  1
  """
  @spec count([char], char) :: non_neg_integer
  def count(strand, nucleotide) do
    do_count(valid_strand?(strand), valid_nucleotide?(nucleotide))
  end

  defp do_count(false, _), do: raise ArgumentError
  defp do_count(_, false), do: raise ArgumentError
  defp do_count(strand, nucleotide) do
    strand
    |> Enum.count(&(&1 == nucleotide))
  end

  defp valid_nucleotide?(nucleotide) when nucleotide in @nucleotides, do: nucleotide
  defp valid_nucleotide?(_), do: false

  defp valid_strand?(strand) do
    cond do
      Enum.all?(strand, &(&1 in 'AGCT')) == true -> strand
      true                                       -> false
    end
  end

  @doc """
  Returns a summary of counts by nucleotide.

  ## Examples

  iex> DNA.nucleotide_counts('AATAA')
  %{?A => 4, ?T => 1, ?C => 0, ?G => 0}
  """
  @spec nucleotide_counts([char]) :: Dict.t
  def nucleotide_counts(strand) do
    @nucleotides
    |> Enum.reduce(%{}, &(Map.put(&2, &1, count(strand, &1))))
  end
end
