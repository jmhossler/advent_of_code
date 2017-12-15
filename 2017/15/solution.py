a_val = 883
b_val = 879

# test
a_val = 65
b_val = 8921

a_factor = 16807
b_factor = 48271
modulo_val = 2147483647

match = 0
iterations = 40000000
for i in range(0, iterations):
    a_val = (a_val * a_factor) % modulo_val
    b_val = (b_val * b_factor) % modulo_val
    a_bits = ('{:64b}'.format(a_val))[-16:]
    b_bits = ('{:64b}'.format(b_val))[-16:]

    if a_bits == b_bits:
        match += 1

print(match)
