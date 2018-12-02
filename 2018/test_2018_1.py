import pytest

from solution1 import get_total_frequency, find_duplicate_frequency


@pytest.mark.parametrize('frequencies, expected', [
    ([+1, -2, +3, +1], 3),
    ([+1, +1, -2], 0),
    ([+1, +1, +1], 3),
    ([-1, -2, -3], -6),
    ])
def test_frequency_analyzer(frequencies, expected):
    assert get_total_frequency(frequencies) == expected

@pytest.mark.parametrize('frequencies, expected', [
    ([+1, -2, +3, +1], 2),
    ([+1, -1], 0),
    ([+3, +3, +4, -2, -4], 10),
    ([-6, +3, +8, +5, -6], 5),
    ([+7, +7, -2, -7, -4], 14),
    ])
def test_frequency_duplicate(frequencies, expected):
    assert find_duplicate_frequency(frequencies) == expected
