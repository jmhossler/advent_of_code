import sys
with open(sys.argv[1], 'r') as f: data = f.read().split('\n')[:-1]

class Particle:
    def __init__(self, position, velocity, acceleration):
        self.position = position
        self.velocity = velocity
        self.acceleration = acceleration


def parse_particle(info):
    print(info.replace(',', '').split(' '))
    return ''

particles = []
for line in data:
    particles.append(parse_particle(line))
