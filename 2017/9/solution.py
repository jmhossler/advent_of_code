import sys
with open(sys.argv[1], 'r') as f: data = f.readline()

def remove_cancelled(data):
    clean_data = []
    i = 0
    while i < len(data):
        if data[i] is '!':
            i += 1
        else:
            clean_data.append(data[i])
        i += 1
    return ''.join(clean_data)

def remove_garbage(data):
    clean_data = []
    is_clean = True
    total = 0
    for val in data:
        if is_clean == True:
            if val is '<':
                is_clean = False
            else:
                clean_data.append(val)
        else:
            if val is '>':
                is_clean = True
            else:
                total += 1
    print('non <> garbage: {}'.format(total))
    return ''.join(clean_data)



data = remove_cancelled(data)
data = remove_garbage(data)
clean_data = []
i = 0
while i < len(data):
    if data[i] in ['{', '}']:
        clean_data.append(data[i])
    i += 1

total = 0
depth = 0
for val in clean_data:
    if val == '}':
        total += depth
        depth -= 1
    if val == '{':
        depth += 1
print(total)
