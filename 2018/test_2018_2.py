import pytest

from solution2 import get_checksum, get_common_letters, solution


@pytest.mark.parametrize('inputs, checksum', [
    (['abcdef', 'bababc', 'abbcde', 'abcccd', 'aabcdd', 'abcdee', 'ababab'], 12),
    (['abcdef'], 0),
    ])
def test_get_checksum(inputs, checksum):
    assert get_checksum(inputs) == checksum

def test_get_common_letters():
    inputs = ['abcde', 'fghij', 'klmno', 'pqrst', 'fguij', 'axcye', 'wvxyz']
    assert get_common_letters(inputs) == 'fgij'

def test_solution():
    assert solution('2018/input/input2.txt') == (7776, 'wlkigsqyfecjqqmnxaktdrhbz')
