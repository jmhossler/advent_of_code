def sum_metadata(numbers):
    metadata_entries = numbers[1]
    children = numbers[0]
    index = 2
    metadata = numbers[-metadata_entries:]
    for _ in range(0, children):
        child_length = get_length(numbers[index:])
        metadata.append(sum_metadata(numbers[index:index+child_length]))
        index += child_length
    return sum(metadata)

def get_length(numbers):
    children = numbers[0]
    metadata_entries = numbers[1]
    index = 2
    children_length = 0
    for _ in range(0, children):
        child_length = get_length(numbers[index:])
        children_length += child_length
        index += child_length
    return children_length + 2 + metadata_entries

def get_value_of_node(numbers):
    metadata_entries = numbers[1]
    children = numbers[0]
    metadata = numbers[-metadata_entries:]
    if children == 0:
        return sum(metadata)
    total_value = 0
    children_entries = []
    index = 2
    for _ in range(0, children):
        child_length = get_length(numbers[index:])
        children_entries.append(numbers[index:index+child_length])
        index += child_length
    for value in metadata:
        if value != 0 and value <= children:
            total_value += get_value_of_node(children_entries[value-1])
    return total_value

def solution(filename):
    with open(filename) as f:
        numbers = [int(x) for x in f.read().strip().split(' ')]
    return sum_metadata(numbers), get_value_of_node(numbers)
