import sys

with open(sys.argv[1], 'r') as f: steps = int(f.read())

buf = [0]
curr_index = 0

for i in range(1, 50000001):
    curr_index = (curr_index + steps) % len(buf)
    next_val = (curr_index + 1)
    buf = buf[:next_val] + [i] + buf[next_val:]
    curr_index = next_val

print(buf[curr_index+1])
print(buf[buf.index(0)+1])
