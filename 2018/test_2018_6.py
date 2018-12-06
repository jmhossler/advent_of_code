from solution6 import get_largest_finite_area, solution, get_largest_region_within_range

POINTS = ['1, 1', '1, 6', '8, 3', '3, 4', '5, 5', '8, 9']

def test_get_largest_finite_area():
    assert get_largest_finite_area(POINTS) == 17

def test_get_largest_region_within_range():
    assert get_largest_region_within_range(POINTS, 32) == 16

def test_solution():
    assert solution('2018/input/input6.txt') == (3722, 44634)
