defmodule Meetup do
  @moduledoc """
  Calculate meetup dates.
  """

  @type weekday ::
      :monday | :tuesday | :wednesday
    | :thursday | :friday | :saturday | :sunday

  @type schedule :: :first | :second | :third | :fourth | :last | :teenth

  @dofw %{monday: 1,
          tuesday: 2,
          wednesday: 3,
          thursday: 4,
          friday: 5,
          saturday: 6,
          sunday: 7}

  @schedule_start %{first: 1,
                    second: 8,
                    third: 15,
                    fourth: 22,
                    teenth: 13}

  @doc """
  Calculate a meetup date.

  The schedule is in which week (1..4, last or "teenth") the meetup date should
  fall.
  """
  @spec meetup(pos_integer, pos_integer, weekday, schedule) :: :calendar.date
  def meetup(year, month, weekday, :last) do
    start_date = :calendar.last_day_of_the_month(year, month) - 6
    _meetup(year, month, @dofw[weekday], start_date)
  end
  def meetup(year, month, weekday, schedule) do
    _meetup(year, month, @dofw[weekday], @schedule_start[schedule])
  end

  defp _meetup(y, m, day, start) do
    cond do
      :calendar.day_of_the_week({y, m, start}) != day ->
        _meetup(y, m, day, start + 1)
      :calendar.day_of_the_week({y, m, start}) == day ->
        {y, m, start}
    end
  end
end

