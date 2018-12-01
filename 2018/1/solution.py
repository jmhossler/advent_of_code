import sys

def frequency_total(frequencies):
    return sum([int(frequency) for frequency in frequencies])

def frequency_duplicate(frequencies):
    previous_frequencies = []
    for frequency in _generate_frequencies(frequencies):
        if frequency in previous_frequencies:
            return frequency
        previous_frequencies.append(frequency)

def _generate_frequencies(frequencies):
    frequency = 0
    index = 0
    while True:
        yield frequency
        frequency += int(frequencies[index])
        index = (index + 1) % len(frequencies)

if __name__ == '__main__':
    with open(sys.argv[1], 'r') as f:
        frequencies = f.readlines()
    print(frequency_total(frequencies))
    print(frequency_duplicate(frequencies))
