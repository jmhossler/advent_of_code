import pytest

from solution import frequency_total, frequency_duplicate


@pytest.mark.parametrize('frequencies, expected', [
    (['+1', '-2', '+3', '+1'], 3),
    (['+1', '+1', '-2'], 0),
    (['+1', '+1', '+1'], 3),
    (['-1', '-2', '-3'], -6),
    ])
def test_frequency_analyzer(frequencies, expected):
    assert frequency_total(frequencies) == expected

@pytest.mark.parametrize('frequencies, expected', [
    (['+1', '-2', '+3', '+1'], 2),
    (['+1', '-1'], 0),
    (['+3', '+3', '+4', '-2', '-4'], 10),
    (['-6', '+3', '+8', '+5', '-6'], 5),
    (['+7', '+7', '-2', '-7', '-4'], 14),
    ])
def test_frequency_duplicate(frequencies, expected):
    assert frequency_duplicate(frequencies) == expected
