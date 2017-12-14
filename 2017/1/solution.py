with open('input', 'r') as f: data = f.read().strip()

total_sum = 0
for i in range(1, len(data)):
    if data[i] == data[i-1]:
        total_sum += int(data[i])

if data[0] == data[len(data)-1]:
    total_sum += int(data[0])

print(total_sum)

new_sum = 0
forward_jump = len(data)/2
for i in range(0, len(data)):
    comparison = int((i + forward_jump) % len(data))
    if data[i] == data[comparison]:
        new_sum += int(data[i])


print(new_sum)
