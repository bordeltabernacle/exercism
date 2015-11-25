defmodule ListOps do
  # Please don't use any external modules (especially List) in your
  # implementation. The point of this exercise is to create these basic functions
  # yourself.
  #
  # Note that `++` is a function from an external module (Kernel, which is
  # automatically imported) and so shouldn't be used either.

  @spec count(list) :: non_neg_integer
  def count([]),    do: 0
  def count([_|t]), do: count(t) + 1

  @spec reverse(list) :: list
  def reverse(l),              do: reverse(l, [])
  defp reverse([], result),    do: result
  defp reverse([h|t], result), do: reverse(t, [h|result])

  @spec map(list, (any -> any)) :: list
  def map(l, f),              do: map(l, f, [])
  defp map([], _, result),    do: reverse(result)
  defp map([h|t], f, result), do: map(t, f, [f.(h)|result])

  @spec filter(list, (any -> as_boolean(term))) :: list
  def filter(l, f),             do: filter(l, f, [])
  defp filter([], _, result),   do: reverse(result)
  defp filter([h|t], f, result) do
    case f.(h) do
      true  -> filter(t, f, [h|result])
      false -> filter(t, f, result)
    end
  end

  @type acc :: any
  @spec reduce(list, acc, ((any, acc) -> acc)) :: acc
  def reduce([], acc, _),       do: acc
  def reduce([h|t], acc, f),    do: reduce(t, f.(h, acc), f)

  @spec append(list, list) :: list
  def append(a, b),              do: append(a, b, reverse(a))
  defp append(_, [], result),    do: reverse(result)
  defp append(a, [h|t], result), do: append(a, t, [h|result])

  @spec concat([[any]]) :: [any]
  def concat(ll),             do: concat(reverse(ll), [])
  defp concat([], result),    do: result
  defp concat([h|t], result), do: concat(t, append(h, result))

end
