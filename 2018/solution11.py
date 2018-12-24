from itertools import product
import numpy as np


def get_power_level(coordinate, serial_number):
    if(coordinate[0] == 0 or coordinate[1] == 0):
        return 0
    rack_id = coordinate[0] + 10
    power_level = coordinate[1] * rack_id
    power_level += serial_number
    power_level *= rack_id
    power_level = int('{0:03d}'.format(power_level)[-3])
    return power_level - 5

def find_largest_total_power(serial_number, square_size):
    a, _ = find_largest_total_power_foo(serial_number, square_size)
    return a

def find_largest_total_power_foo(serial_number, square_size):
    points = product(range(1, 301 - square_size), range(1, 301 - square_size))
    board = []
    for x in range(0, 301):
        board.append([])
        for y in range(0, 301):
            board[x].append(0)
            board[x][y] = get_power_level((x, y), serial_number)
    b = np.cumsum(np.array(board), axis=0)
    c = np.cumsum(b, axis=1)
    max_point = (1, 1)
    max_sum = get_total(c, max_point, square_size)
    for point in points:
        square_sum = get_total(c, point, square_size)
        if square_sum > max_sum:
            max_point = point
            max_sum = square_sum
    return max_point, max_sum

def find_largest_total_power_x(serial_number):
    largest_point, largest_sum = find_largest_total_power_foo(serial_number, 1)
    largest_size = 1
    for x in range(1, 301):
        a, b = find_largest_total_power_foo(serial_number, x)
        if b > largest_sum:
            largest_point, largest_sum = a, b
            largest_size = x
    return largest_point, largest_size

def get_total(cumulative_sum, point, square_size):
    a, b, c, d = ((point[0] - 1 + square_size, point[1] - 1 + square_size),
                  (point[0] - 1, point[1] - 1),
                  (point[0] - 1 + square_size, point[1] - 1),
                  (point[0] - 1, point[1] - 1 + square_size),
                 )
    return (cumulative_sum[a[0]][a[1]]
            + cumulative_sum[b[0]][b[1]]
            - cumulative_sum[c[0]][c[1]]
            - cumulative_sum[d[0]][d[1]])

def create_cumulative_sum(board):
    new_board = []
    for x in range(0, len(board)):
        new_board.append([])
        for y in range(0, len(board[0])):
            new_board[x].append(0)
            new_board[x][y] = get_sum(board, x, y)
    return new_board

def get_sum(board, x, y):
    total_sum = 0
    for i in range(0, x+1):
        for j in range(0, y+1):
            total_sum += board[i][j]
    return total_sum

if __name__ == '__main__':
    find_largest_total_power(18, 3)
