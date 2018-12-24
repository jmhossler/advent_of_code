import pytest

from solution11 import get_power_level, find_largest_total_power, find_largest_total_power_x

@pytest.mark.parametrize(
        'coordinate,serial_number,expected',
        [((3, 5), 8, 4),
         ((122, 79), 57, -5),
         ((217, 196), 39, 0),
         ((101, 153), 71, 4),
        ])
def test_get_power_level(coordinate, serial_number, expected):
    assert get_power_level(coordinate, serial_number) == expected

@pytest.mark.parametrize(
        'serial,expected',
        [(18, (33, 45)),
         (42, (21, 61)),
        ])
def test_find_largest_total_power(serial, expected):
    assert find_largest_total_power(serial, 3) == expected

def test_find_largest_total_power_x():
    assert find_largest_total_power_x(18) == ((90, 269), 16)
    assert find_largest_total_power_x(42) == ((232, 251), 12)

def test_solution():
    assert find_largest_total_power_x(8199) == ((234, 272), 18)
