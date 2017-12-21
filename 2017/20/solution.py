import sys
import re
with open(sys.argv[1], 'r') as f: data = f.read().split('\n')[:-1]

class Particle:
    def __init__(self, position, velocity, acceleration, id):
        self.id = id
        self.position = position
        self.velocity = velocity
        self.acceleration = acceleration

    def step(self):
        for i in range(0, 3):
            self.velocity[i] += self.acceleration[i]
            self.position[i] += self.velocity[i]


def parse_particle(info, i):
    data = info.split(' ')
    pattern = '<(.*)>'
    p = [int(x) for x in re.search(pattern, data[0]).group(1).split(',')]
    v = [int(x) for x in re.search(pattern, data[1]).group(1).split(',')]
    a = [int(x) for x in re.search(pattern, data[2]).group(1).split(',')]
    return Particle(p, v, a, i)

particles = []
for i in range(0, len(data)):
    particles.append(parse_particle(data[i], i))

def get_abs(triplet):
    return sum([x ** 2 for x in triplet]) ** (1/2)

def get_man(triplet):
    return sum([abs(x) for x in triplet])

particles = sorted(particles, key=lambda x: get_abs(x.acceleration))

print(particles[0].id)

for i in range(0, 1000):
    positions = [tuple(x.position) for x in particles]
    if len(positions) != len(set(positions)):
        for position in positions:
            if positions.count(position) > 1:
                particles = [particle for particle in particles if tuple(particle.position) != position]
    for particle in particles:
        particle.step()
print(len(particles))
