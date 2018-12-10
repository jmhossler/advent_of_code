import pytest

from solution9 import get_high_score, solution

@pytest.mark.parametrize(
        'players,marbles,expected',
        [(5, 25, 32),
         (10, 1618, 8317),
         (13, 7999, 146373),
         (17, 1104, 2764),
         (21, 6111, 54718),
         (30, 5807, 37305),
         ])
def test_get_high_score(players, marbles, expected):
    assert get_high_score(players, marbles) == expected

def test_solution():
    assert solution('2018/input/input9.txt') == (375465, None)
