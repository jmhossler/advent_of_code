import pytest

from solution5 import break_down_polymer, improve_polymer, solution


@pytest.mark.parametrize('code,expected', [
    ('aA', ''),
    ('abBA', ''),
    ('abAB', 'abAB'),
    ('aabAAB', 'aabAAB'),
    ('dabAcCaCBAcCcaDA', 'dabCBAcaDA'),
    ])
def test_break_down_polymer(code, expected):
    assert break_down_polymer(code) == expected

def test_improve_polymer():
    assert improve_polymer('dabAcCaCBAcCcaDA') == 'daDA'

def test_solution():
    a, b = solution('2018/input/input5.txt')
    assert (len(a), len(b)) == (10384, 5412)
