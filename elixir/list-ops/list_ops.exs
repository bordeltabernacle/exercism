defmodule ListOps do

  @moduledoc """
  Implements basic list operations

  In functional languages list operations like `length`, `map`, and
  `reduce` are very common.
  Implement a series of basic list operations,
  without using existing functions.
  """

  @doc """
  Returns the collection’s size.
  """
  @spec count(list) :: non_neg_integer
  def count([]),       do: 0
  def count([_|tail]), do: count(tail) + 1

  @doc """
  Reverses the collection.
  """
  @spec reverse(list) :: list
  def reverse(l),                    do: reverse(l, [])

  defp reverse([], result),          do: result
  defp reverse([head|tail], result), do: reverse(tail, [head|result])

  @doc """
  Returns a new collection, where each item is the result of
  invoking `func` on each corresponding item of the collection.
  """
  @spec map(list, (any -> any)) :: list
  def map(l, func),                   do: map(l, func, [])

  defp map([], _, result),            do: reverse(result)
  defp map([head|tail], func, result) do
    map(tail, func, [func.(head)|result])
  end

  @doc """
  Filters the collection, i.e. returns only those elements for
  which `func` returns a truthy value.
  """
  @spec filter(list, (any -> as_boolean(term))) :: list
  def filter(l, func),                   do: filter(l, func, [])

  defp filter([], _, result),            do: reverse(result)
  defp filter([head|tail], func, result) do
    case func.(head) do
      true  -> filter(tail, func, [head|result])
      false -> filter(tail, func, result)
    end
  end

  @doc """
  Invokes `func` for each element in the collection passing that
  element and the accumulator `acc` as arguments.
  `func`‘s return value is stored in `acc`. The first element of
  the collection is used as the initial value of `acc`.

  Note that since the first element of the enumerable is used as
  the initial value of the accumulator, `func` will only be
  executed `n - 1` times where `n` is the length of the enumerable.
  """
  @type acc :: any
  @spec reduce(list, acc, ((any, acc) -> acc)) :: acc
  def reduce([], acc, _),            do: acc
  def reduce([head|tail], acc, func) do
    reduce(tail, func.(head, acc), func)
  end

  @doc """
  Inserts the second collection into the end of the first collection.

  Returns a new collection which contains the elements from the first
  collection followed by the elements from the second collection.
  """
  @spec append(list, list) :: list
  def append(a, b),              do: append(a, b, reverse(a))

  defp append(_, [], result),    do: reverse(result)
  defp append(a, [head|tail], result) do
    append(a, tail, [head|result])
  end

  @doc """
  Given an enumerable of enumerables, concatenates the enumerables
  into a single list.
  """
  @spec concat([[any]]) :: [any]
  def concat(ll),             do: concat(reverse(ll), [])

  defp concat([], result),    do: result
  defp concat([head|tail], result) do
    concat(tail, append(head, result))
  end

end
