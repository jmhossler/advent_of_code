import re


def get_states(raw_light_data, time):
    lights = parse_raw_light_data(raw_light_data)
    positions = [x['position'] for x in move_lights(lights, 10144)]
    x_vals = {x[0] for x in positions}
    y_vals = {x[1] for x in positions}
    min_x = min(x_vals)
    min_y = min(y_vals)
    max_x = max(x_vals)
    max_y = max(y_vals)
    foo = {y: set({}) for y in y_vals}
    for position in positions:
        foo[position[1]].add(position[0])
    lines = []
    for y in range(min_y, max_y+1):
        line = ['.' for x in range(min(x_vals), max(x_vals)+1)]
        if y in foo:
            for x in foo[y]:
                line[x-min(x_vals)] = '#'
        lines.append(''.join(line))
    for line in lines:
        print(''.join(line))

def move_lights(lights, seconds):
    new_lights = []
    for light in lights:
        new_lights.append(move_light(light, seconds))
    return new_lights

def move_light(light, seconds):
    position_x = light['position'][0]
    position_y = light['position'][1]
    delta_x = light['velocity'][0]
    delta_y = light['velocity'][1]
    return {'position': (position_x + (delta_x * seconds), position_y + (delta_y * seconds)), 'velocity': light['velocity']}

def parse_raw_light_data(light_data):
    lights = []
    for light in light_data:
        lights.append({'position': parse_position(light), 'velocity': parse_velocity(light)})
    return lights

def parse_position(light):
    position = re.search('<.*?>', light).group()[1:-1].split(',')
    return (int(position[0]), int(position[1]))

def parse_velocity(light):
    velocity = re.findall('<.*?>', light)[1][1:-1].split(',')
    return (int(velocity[0]), int(velocity[1]))

if __name__ == '__main__':
    with open('2018/input/input10.txt') as f:
        light_data = f.readlines()
    boards = get_states(light_data, 100000)
