b_val = 879

# test
#b_val = 8921

a_factor = 16807
b_factor = 48271
modulo_val = 2147483647

def a_generator(n):
    a_val = 883
    count = 0
    while count < n:
        a_val = int((a_val * a_factor) % modulo_val)
        if a_val % 4 == 0:
            yield a_val
            count += 1

def b_generator(n):
    b_val = 879
    count = 0
    while count < n:
        b_val = int((b_val * b_factor) % modulo_val)
        if b_val % 8 == 0:
            yield b_val
            count += 1


match = 0
iterations = 5000000
for a_val, b_val in zip(a_generator(iterations), b_generator(iterations)):
    a_bits = ('{:64b}'.format(a_val))[-16:]
    b_bits = ('{:64b}'.format(b_val))[-16:]
    if a_bits == b_bits:
        match += 1

print(match)
