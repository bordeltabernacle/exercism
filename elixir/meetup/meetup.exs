defmodule Meetup do
  @moduledoc """
  Calculate meetup dates.
  """

  @type weekday ::
      :monday | :tuesday | :wednesday
    | :thursday | :friday | :saturday | :sunday

  @type schedule :: :first | :second | :third | :fourth | :last | :teenth

  @doc """
  Calculate a meetup date.

  The schedule is in which week (1..4, last or "teenth") the meetup date should
  fall.
  """
  @spec meetup(pos_integer, pos_integer, weekday, schedule) :: :calendar.date
  def meetup(year, month, weekday, :last) do
    start_date = :calendar.last_day_of_the_month(year, month) - 6
    _meetup(year, month, dofw(weekday), start_date)
  end
  def meetup(year, month, weekday, schedule) do
    _meetup(year, month, dofw(weekday), start_date(schedule))
  end

  defp _meetup(y, m, day, start) do
    cond do
      :calendar.day_of_the_week({y, m, start}) != day ->
        _meetup(y, m, day, start + 1)
      :calendar.day_of_the_week({y, m, start}) == day ->
        {y, m, start}
    end
  end

  defp dofw(:monday),    do: 1
  defp dofw(:tuesday),   do: 2
  defp dofw(:wednesday), do: 3
  defp dofw(:thursday),  do: 4
  defp dofw(:friday),    do: 5
  defp dofw(:saturday),  do: 6
  defp dofw(:sunday),    do: 7

  defp start_date(:first),  do: 1
  defp start_date(:second), do: 8
  defp start_date(:third),  do: 15
  defp start_date(:fourth), do: 22
  defp start_date(:teenth), do: 13
end

