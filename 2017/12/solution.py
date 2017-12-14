data = {}
with open('input', 'r') as f:
    for line in f:
        values = line.split()
        data[values[0]] = [x.replace(',', '') for x in values[2:]]

def get_connected(data, value):
    values = [value] + data[value]
    new_values = data[value]
    while len(new_values) > 0:
        temp_new_values = []
        for val in new_values:
            for other_val in data[val]:
                if other_val not in values and other_val not in new_values and other_val not in temp_new_values:
                    temp_new_values.append(other_val)
        values += new_values
        new_values = temp_new_values
    return list(set(values))

print(len(get_connected(data, '0')))

groups = 0
in_groups = []
for val in data:
    if val not in in_groups:
        in_groups += get_connected(data, val)
        groups += 1
print(groups)
