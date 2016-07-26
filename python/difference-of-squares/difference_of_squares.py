square_of_sum = lambda n: ((n * (n + 1)) / 2)**2
sum_of_squares = lambda n: (n * (n + 1) * (2 * n + 1)) / 6
difference = lambda n: square_of_sum(n) - sum_of_squares(n)
