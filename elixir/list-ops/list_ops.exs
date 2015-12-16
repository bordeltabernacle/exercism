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
  def count(list) do
    list
    |> reduce 0, fn(_, acc) -> acc + 1 end
  end

  @doc """
  Reverses the collection.
  """
  @spec reverse(list) :: list
  def reverse(list) do
    list
    |> reduce [], fn(elem, acc) -> [elem|acc] end
  end

  @doc """
  Returns a new collection, where each item is the result of
  invoking `func` on each corresponding item of the collection.
  """
  @spec map(list, (any -> any)) :: list
  def map(list, func) do
    list
    |> reverse
    |> reduce [], fn(elem, acc) -> [func.(elem)|acc] end
  end

  @doc """
  Filters the collection, i.e. returns only those elements for
  which `func` returns a truthy value.
  """
  @spec filter(list, (any -> as_boolean(term))) :: list
  def filter(list, func) do
    list
    |> reverse
    |> reduce [], fn(elem, acc) ->
      case func.(elem) do
        true  -> [elem|acc]
        false -> acc
      end
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
  def reduce([], acc, _func), do: acc
  def reduce([head|tail], acc, func) do
    reduce(tail, func.(head, acc), func)
  end

  @doc """
  Inserts the second collection into the end of the first collection.

  Returns a new collection which contains the elements from the first
  collection followed by the elements from the second collection.
  """
  @spec append(list, list) :: list
  def append(list_a, list_b) do
    list_a
    |> reverse
    |> reduce list_b, fn(elem, acc) -> [elem|acc] end
  end

  @doc """
  Given an enumerable of enumerables, concatenates the enumerables
  into a single list.
  """
  @spec concat([[any]]) :: [any]
  def concat(list_of_lists) do
    list_of_lists
    |> reverse
    |> reduce [], fn(elem, acc) -> append elem, acc end
  end

end
