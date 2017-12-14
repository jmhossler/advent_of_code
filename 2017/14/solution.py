import sys

def reverse_in_place(array, start, stop):
    sub_array = get_sub_array(array, start, stop)
    sub_array.reverse()
    for i in range(0, len(sub_array)):
        array[(start + i) % len(array)] = sub_array[i]

def get_sub_array(array, start, stop):
    values = []
    i = start
    while i != (stop % len(array)):
        values.append(array[i])
        i = (i + 1) % len(array)
    return values

def swap(array, i, j):
    temp = array[i]
    array[i] = array[j]
    array[j] = temp

def knot_hash(data):
    data_array = [i for i in range(0, 256)]

    lengths = [ord(x) for x in data]
    lengths += [17, 31, 73, 47, 23]

    skip_value = 0
    index = 0

    for _ in range(0, 64):
        for length in lengths:
            reverse_in_place(data_array, index, index + length)
            index = (index + length + skip_value) % len(data_array)
            skip_value += 1
    boxes = [data_array[i:i+16] for i in range(0, len(data_array), 16)]
    output_str = ''
    for box in boxes:
        xor_val = 0
        for val in box:
            xor_val ^= val
        output_str += '{:02x}'.format(xor_val)
    return output_str

def convert_hash(hash_str):
    ret_str = ''
    for val in hash_str:
        val = int(val, 16)
        ret_str += '{:04b}'.format(val)
    return ret_str

with open(sys.argv[1], 'r') as f: data = f.read()

hash_value = data.replace('\n', '')
hex_grid = []
total_count = 0
for i in range(0, 128):
    print(i)
    row_hash = hash_value + '-' + str(i)
    row = convert_hash(knot_hash(row_hash))
    hex_grid.append(row)
    total_count += row.count('1')

print(total_count)

def surrounding(grid, coord):
    for i in range(-1, 2, 2):
        for j in range(-1, 2, 2):
            yield (i, j)

def get_adjacent(grid, coord):
    adjacent = []
    for other in surrounding(grid, coord):
        if grid[other[0]][other[1]] == '1':
            adjacent.append(other)
    return adjacent

adjacents = []
regions = 0
for i in range(0, len(hex_grid)):
    for j in range(0, len(hex_grid[i])):
        print(adjacents)
        if hex_grid[i][j] == '1' and (i, j) not in adjacents:
            regions += 1
            queue = [(i, j)]
            while len(queue) != 0:
                curr = queue.pop(0)
                adjacents.append(curr)
                for val in get_adjacent(hex_grid, curr):
                    if val not in queue:
                        queue.append(val)

print(regions)
