import sys

with open(sys.argv[1], 'r') as f: data = f.read().split('\n')[:-1]

rules = {}
for line in data:
    line = line.split(' ')
    rules[line[0]] = line[2]

start = ".#./..#/###"

def count_on(grid):
    return grid.count('#')

def flip_h(grid):
    sub = grid.split('/')
    return '/'.join(sub[::-1])

def flip_v(grid):
    sub = grid.split('/')
    return '/'.join([x[::-1] for x in sub])

def rotate(grid):
    sub = grid.split('/')
    new_vals = [[None for x in sub] for x in sub]
    new_val_map = {(x, y): (y, len(sub) - 1 - x) for x in range(0, len(sub)) for y in range(0, len(sub))}
    for coord in new_val_map:
        new_coord = new_val_map[coord]
        new_vals[new_coord[0]][new_coord[1]] = sub[coord[0]][coord[1]]
    return '/'.join([''.join(x) for x in new_vals])


def pretty_grid(grid):
    sub = grid.split('/')
    return '\n'.join(sub)

def join_grids(grids, size):
    if len(grids) == 1:
        return grids[0]
    quadrants = [grids[i:i+size] for i in range(0, len(grids), size)]
    upper = grids[:int(len(grids)/2)]
    lower = grids[int(len(grids)/2):]
    rows = []
    for quad in quadrants:
        y = [x.split('/') for x in quad]
        for i in range(0, len(y[0])):
            rows.append(''.join([x[i] for x in y]))
    return '/'.join(rows)

def get_smaller(grid, size):
    grid_expanded = grid.split('/')
    grids = []
    rows = [grid_expanded[i:i+size]
            for i in range(0, len(grid_expanded), size)]
    for row in rows:
        array = [x[i:i+size]
                 for i in range(0, len(grid_expanded), size)
                 for x in row]
        for val in [array[i:i+size] for i in range(0, len(array), size)]:
            grids.append('/'.join(val))
    return grids

def expand(grid):
    grid_expanded = grid.split('/')
    smaller = []
    new_grid_len = 0
    size = 0
    if len(grid_expanded) % 2 == 0:
        size = 2
        new_grid_len = 3
        smaller = get_smaller(grid, 2)
    else:
        size = 3
        new_grid_len = 4
        smaller = get_smaller(grid, 3)
    new_grids = []
    for sub_grid in smaller:
        mirror = flip_v(sub_grid)
        while mirror not in rules and sub_grid not in rules:
            mirror = rotate(mirror)
            sub_grid = rotate(sub_grid)
        if mirror in rules:
            new_grids.append(rules[mirror])
        else:
            new_grids.append(rules[sub_grid])
    if len(new_grids) == 1:
        return join_grids(new_grids, new_grid_len / 2)
    else:
        return join_grids(new_grids, int(len(grid_expanded) / size))


grid = start
for i in range(0, 18):
    grid = expand(grid)

print(count_on(grid))
