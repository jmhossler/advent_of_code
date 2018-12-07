from solution7 import get_step_order, solution, get_step_time


INPUT = [
    "Step C must be finished before step A can begin.",
    "Step C must be finished before step F can begin.",
    "Step A must be finished before step B can begin.",
    "Step A must be finished before step D can begin.",
    "Step B must be finished before step E can begin.",
    "Step D must be finished before step E can begin.",
    "Step F must be finished before step E can begin.",
]

def test_get_step_order():
    assert get_step_order(INPUT) == 'CABDFE'

def test_get_step_time():
    assert get_step_time(INPUT, 2, 0) == 15

def test_solution():
    assert solution('2018/input/input7.txt') == ("JKNSTHCBGRVDXWAYFOQLMPZIUE", 755)
