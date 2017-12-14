import sys

def get_floor(line):
    return int(line.split()[0][:-1])

def get_range(line):
    return int(line.split()[1])

def update_floors(floors):
    new_floors = {}
    for floor in floors:
        direction = floors[floor]['direction']
        if floors[floor]['place'] == 1:
            direction = 1
        elif floors[floor]['place'] == floors[floor]['max']:
            direction = -1
        new_floors[floor] = {'place': floors[floor]['place'] + direction,
                             'max': floors[floor]['max'],
                             'direction': direction}
    return new_floors

def display_floors(floors):
    display_str = ''
    for floor in floors:
        floor_str = ['.'] * floors[floor]['max']
        floor_str[floors[floor]['place'] - 1] = 'x'
        display_str += '{}: {}\n'.format(floor, ''.join(floor_str))
    return display_str

def get_copy_of_floors(floors):
    return {val: {'place': floors[val]['place'], 'max': floors[val]['max'], 'direction': floors[val]['direction']}
            for val in floors}

floors = {}
with open(sys.argv[1], 'r') as f:
    for line in f:
        floors[get_floor(line)] = {'place': 1, 'max': get_range(line), 'direction': 1}

pt1_floors = get_copy_of_floors(floors)

severity = 0
for i in range(0, max(pt1_floors.keys()) + 1):
    if i in pt1_floors:
        if pt1_floors[i]['place'] == 1:
            severity += i * pt1_floors[i]['max']
    pt1_floors = update_floors(pt1_floors)

print(severity)

def update_floors_n_times(floors, n):
    new_floors = {}
    for floor in floors:
        direction = floors[floor]['direction']
        x = n % ((floors[floor]['max'] * 2) - 2)
        placements = [x for x in range(1, floors[floor]['max'] + 1)]
        sub_vals = placements[:-1]
        sub_vals.reverse()
        placements += sub_vals[:-1]
        x = (x + floors[floor]['place']) % ((floors[floor]['max'] * 2) - 2)
        placement = placements[x]
        if x < floors[floor]['max']:
            direction = 1
        else:
            direction = -1
        new_floors[floor] = {'place': placement,
                             'max': floors[floor]['max'],
                             'direction': direction}
    return new_floors

def detected(floors, start_time):
    print('starting with {}'.format(start_time))
    floors = get_copy_of_floors(floors)
    print('plain floors, default')
    print(display_floors(floors))
    floors = update_floors_n_times(floors, start_time)
    print('floors shifted by {}'.format(start_time))
    print(display_floors(floors))

    last_floor = 0
    security_floors = floors.keys()
    for floor in security_floors:
        print(floor)
        if floor != last_floor:
            print(last_floor)
            floors = update_floors_n_times(floors, floor - last_floor)
        print(display_floors(floors))
        if floors[floor]['place'] == 1:
            return True
        last_floor = floor

    # for i in range(0, max(floors.keys()) + 1):
    #     if i in floors:
    #         if floors[i]['place'] == 1:
    #             return True
    #     floors = update_floors(floors)
    return False

pt2_floors = get_copy_of_floors(floors)

pico_seconds = 0
while detected(pt2_floors, pico_seconds) and pico_seconds < 11:
    pico_seconds += 1
print(pico_seconds)
