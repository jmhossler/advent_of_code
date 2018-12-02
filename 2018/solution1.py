import sys
from itertools import accumulate, cycle

from common import read_numbers


def get_total_frequency(frequency_deltas):
    return sum(frequency_deltas)


def find_duplicate_frequency(frequency_deltas):
    previous_frequencies = {0}
    return next(frequency
                for frequency in accumulate(cycle(frequency_deltas))
                if frequency in previous_frequencies or previous_frequencies.add(frequency))


def solution(file_name):
    FREQUENCY_DELTAS = read_numbers(file_name)
    return (get_total_frequency(FREQUENCY_DELTAS),
            find_duplicate_frequency(FREQUENCY_DELTAS))


if __name__ == '__main__':
    print(solution(sys.argv[1]))
