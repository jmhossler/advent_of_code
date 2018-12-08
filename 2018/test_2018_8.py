import pytest
from solution8 import sum_metadata, solution, get_value_of_node

EXAMPLE_A = [2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2]
INPUT = [(EXAMPLE_A, 138),
         ([0, 3, 1, 2, 3], 6),
         ]

@pytest.mark.parametrize('numbers, expected', INPUT)
def test_sum_metadata(numbers, expected):
    assert sum_metadata(numbers) == expected

def test_get_value_of_node():
    assert get_value_of_node(EXAMPLE_A) == 66

def test_solution():
    assert solution('2018/input/input8.txt') == (40848, 34466)
