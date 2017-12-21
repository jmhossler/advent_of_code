import sys

with open(sys.argv[1], 'r') as f: data = f.read()

network_map = [[x for x in line] for line in data.split('\n')[:-1]]

position = (0, network_map[0].index('|'))
direction = (1, 0)
letters = []

def add_positions(a, b):
    return (a[0] + b[0], a[1] + b[1])

def next_position(position, direction):
    possible_next = add_positions(position, direction)
    if get_at(possible_next) == ' ':
        left, next_position = left_position(position, direction)
        if get_at(left) == ' ':
            return right_position(position, direction)
        else:
            return left, next_position
    return possible_next, direction

def get_at(position):
    global network_map
    if ((position[0] < 0 or position[0] >= len(network_map))
        or (position[1] < 0 or position[1] >= len(network_map[position[0]]))):
        return ' '
    return network_map[position[0]][position[1]]

def left_position(position, direction):
    new_direction = (direction[1], direction[0])
    return add_positions(position, new_direction), new_direction

def right_position(position, direction):
    new_direction = (direction[1] * -1, direction[0] * -1)
    return add_positions(position, new_direction), new_direction

def has_next(position, direction):
    next_p, next_d = next_position(position, direction)
    if get_at(next_p) != ' ':
        return True
    return False


steps = 1
while has_next(position, direction):
    steps += 1
    position, direction = next_position(position, direction)
    if get_at(position) not in ['|', '-', '+']:
        letters.append(get_at(position))

print(''.join(letters))
print(steps)
