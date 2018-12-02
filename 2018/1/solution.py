import sys
from itertools import accumulate, cycle

def get_total_frequency(frequency_deltas):
    return sum(frequency_deltas)

def find_duplicate_frequency(frequency_deltas):
    previous_frequencies = {0}
    return next(frequency
                for frequency in accumulate(cycle(frequency_deltas))
                if frequency in previous_frequencies or previous_frequencies.add(frequency))

if __name__ == '__main__':
    with open(sys.argv[1], 'r') as f:
        FREQUENCY_DELTAS = [int(row) for row in f.readlines()]
    print(get_total_frequency(FREQUENCY_DELTAS))
    print(find_duplicate_frequency(FREQUENCY_DELTAS))
