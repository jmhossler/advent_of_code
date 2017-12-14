import sys

def north(location):
    return (location[0], location[1] + 1)

def north_east(location):
    return (location[0] + 0.5, location[1] + 0.5)

def north_west(location):
    return (location[0] - 0.5, location[1] + 0.5)

def south(location):
    return (location[0], location[1] - 1)

def south_east(location):
    return (location[0] + 0.5, location[1] - 0.5)

def south_west(location):
    return (location[0] - 0.5, location[1] - 0.5)

function_map = {
        'n': north,
        'ne': north_east,
        'nw': north_west,
        's': south,
        'se': south_east,
        'sw': south_west,
        }

def distance(a, b):
    return ((b[0] - a[0]) ** 2 + (b[1] - a[1]) ** 2) ** (1 / 2)

def surrounding_locations(location):
    return [north(location),
            north_east(location),
            north_west(location),
            south(location),
            south_east(location),
            south_west(location),
            ]

def get_next(current, final):
    next_place = current
    for location in surrounding_locations(current):
        if distance(location, final) < distance(next_place, final):
            next_place = location
    return next_place

def get_final_location(directions):
    current = (0, 0)
    for direction in directions:
        current = function_map[direction](current)
    return current

def get_furthest_location(directions):
    current = (0, 0)
    furthest_location = (0, 0)
    max_steps = 0
    for direction in directions:
        current = function_map[direction](current)
        current_steps = get_steps(current)
        if current_steps > max_steps:
            furthest_location = current
            max_steps = current_steps
    return max_steps


with open(sys.argv[1], 'r') as f: directions = f.read().replace('\n', '').split(',')
final_location = get_final_location(directions)


def get_steps(final_location):
    current_location = (0, 0)

    steps = 0
    while distance(current_location, final_location) != 0:
        current_location = get_next(current_location, final_location)
        steps += 1

    return steps

def positive_changes(current_location, final_location):
    changes = []
    for location in surrounding_locations(current_location['location']):
        if distance(current_location['location'], final_location) >= distance(location, final_location):
            changes.append({'location': location, 'steps': current_location['steps']})
    return changes


def smart_bfs(final_location):
    location_queue = [{'steps': 0, 'location': (0, 0)}]
    current_location = location_queue.pop(0)
    while distance(current_location['location'], final_location) != 0:
        new_locations = positive_changes(current_location, final_location)
        for location in new_locations:
            if location['location'] not in [val['location'] for val in location_queue]:
                location_queue.append(location)
        current_location = location_queue.pop(0)
    return current_location['steps']

print('Part 1: {}'.format(get_steps(final_location)))
print('Part 2: {}'.format(get_furthest_location(directions)))
#print('Part 2: {}'.format(smart_bfs(furthest_location)))
