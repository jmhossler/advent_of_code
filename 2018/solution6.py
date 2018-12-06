from itertools import product

from common import read_strings


def get_largest_finite_area(points):
    parsed_points = []
    for point in points:
        parsed_points.append(tuple(int(x) for x in point.split(',')))
    max_x = max([point[0] for point in parsed_points]) + 100
    max_y = max([point[1] for point in parsed_points]) + 100
    min_x = min([point[0] for point in parsed_points]) - 100
    min_y = min([point[1] for point in parsed_points]) - 100

    point_map = {point: [] for point in parsed_points}
    for point in product(range(min_x, max_x + 1), range(min_y, max_y + 1)):
        closest_points = sorted(
            parsed_points,
            key=lambda solid_point: manhatten_distance(point, solid_point))
        if (manhatten_distance(point, closest_points[0])
                != manhatten_distance(point, closest_points[1])):
            point_map[closest_points[0]].append(point)
    not_infinite_points = [
        point
        for point in point_map
        if not_infinite(point_map, point, (min_x, max_x), (min_y, max_y))
    ]
    print(not_infinite_points)
    return max(len(point_map[point]) for point in not_infinite_points)

def not_infinite(point_map, point, x, y):
    return not any(a[0] == x[0] or a[0] == x[1] or a[1] == y[0] or a[1] == y[1]
                   for a in point_map[point])

def manhatten_distance(a, b):
    return abs(a[0] - b[0]) + abs(a[1] - b[1])

def get_largest_region_within_range(points, max_range):
    parsed_points = []
    for point in points:
        parsed_points.append(tuple(int(x) for x in point.split(',')))

    max_x = max([point[0] for point in parsed_points])
    max_y = max([point[1] for point in parsed_points])
    min_x = min([point[0] for point in parsed_points])
    min_y = min([point[1] for point in parsed_points])
    x_delta = max_x - min_x
    y_delta = max_y - min_y
    return len([point
                for point in product(
                    range(min_x - max_range + x_delta, max_x + 1 + max_range - x_delta),
                    range(min_y - max_range + y_delta, max_y + 1 + max_range - y_delta))
                if within_range(parsed_points, point, max_range)])

def within_range(points, point, max_range):
    return sum(manhatten_distance(x, point) for x in points) < max_range

def solution(input_file):
    points = read_strings(input_file)
    return (get_largest_finite_area(points), get_largest_region_within_range(points, 10000))
