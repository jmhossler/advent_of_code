from solution3 import parse_input, count_overlap, solution, get_0_overlap_id


INPUT_STRS = [
    "#1 @ 1,3: 4x4",
    "#2 @ 3,1: 4x4",
    "#3 @ 5,5: 2x2",
]


def test_parse_input():
    assert parse_input(INPUT_STRS) == {
        1: [(1, 3), (4, 4)],
        2: [(3, 1), (4, 4)],
        3: [(5, 5), (2, 2)],
    }


def test_find_overlap():
    assert count_overlap(INPUT_STRS) == 4


def test_get_0_overlap_id():
    assert get_0_overlap_id(INPUT_STRS) == 3


def test_solution():
    assert solution('2018/input/input3.txt') == (103806, 625)
