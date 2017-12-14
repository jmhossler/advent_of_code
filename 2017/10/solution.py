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

data_array = [i for i in range(0, 256)]

with open('input', 'r') as f: lengths = [ord(x) for x in f.read().replace('\n', '')]
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
print(output_str)
print(len(output_str))
