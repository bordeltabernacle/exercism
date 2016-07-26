square_of_sum = lambda n: sum(range(n + 1))**2
sum_of_squares = lambda n: sum(i * i for i in range(n + 1))
difference = lambda n: square_of_sum(n) - sum_of_squares(n)
