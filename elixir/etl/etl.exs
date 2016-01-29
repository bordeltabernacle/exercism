defmodule ETL do
  def transform(input) do
    for {k, v} <- input, elem <- v,
    into: %{},
    do: {String.downcase(elem), k}
  end
end
