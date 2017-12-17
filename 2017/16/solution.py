import sys

programs = [chr(ord('a') + i) for i in range(0, 16)]
original = [x for x in programs]

class Node:
    def __init__(self, val, n, p):
        self.val = val
        self.n = n
        self.p = p

    def get_previous(self):
        return self.p

    def set_previous(self, p):
        self.p = p

    def get_next(self):
        return self.n

    def set_next(self, n):
        self.n = n

    def get_val(self):
        return self.val

    def set_val(self, val):
        self.val = val

class LinkedList:
    def __init__(self, values):
        nodes = [Node(val, None, None) for val in values]
        for i in range(0, len(nodes)):
            if i > 0:
                nodes[i].set_previous(nodes[i-1])
            else:
                nodes[i].set_previous(nodes[-1])
            if i < len(nodes) - 1:
                nodes[i].set_next(nodes[i+1])
            else:
                nodes[i].set_next(nodes[0])

        self.head = nodes[0]
        self._quick_access = {node.get_val(): node for node in nodes}

    def spin(self, data):
        spin_size = int(data)
        while spin_size > 0:
            self.head = self.head.get_previous()
            spin_size -= 1

    def swap_i(self, data):
        a_val = int(data.split('/')[0])
        b_val = int(data.split('/')[1])
        probe = self.head
        index = 0
        node_a = None
        node_b = None
        while index <= a_val or index <= b_val:
            if index == a_val:
                node_a = probe
            if index == b_val:
                node_b = probe
            probe = probe.get_next()
            index += 1
        self.swap_n(node_a, node_b)

    def swap_p(self, data):
        a_val = data.split('/')[0]
        b_val = data.split('/')[1]
        nodes = [self._quick_access[node] for node in self._quick_access]
        self._quick_access = {node.get_val(): node for node in nodes}
        node_a = self._quick_access[a_val]
        node_b = self._quick_access[b_val]
        self.swap_n(node_a, node_b)

    def swap_n(self, node_a, node_b):
        node_a_val = node_a.get_val()
        node_a.set_val(node_b.get_val())
        node_b.set_val(node_a_val)

    def join(self):
        ret_str = ''
        probe = self.head
        ret_str += self.head.get_val()
        probe = self.head.get_next()
        while probe != self.head:
            ret_str += probe.get_val()
            probe = probe.get_next()
        return ret_str

    def execute(self, move):
        command = move[0]
        args = move[1:]
        if command == 's':
            self.spin(args)
        elif command == 'x':
            self.swap_i(args)
        elif command == 'p':
            self.swap_p(args)


def spin(programs, data):
    spin_size = int(data)
    return programs[len(programs) - spin_size:] + programs[:len(programs) - spin_size]

def swap_i(programs, data):
    a_val = int(data.split('/')[0])
    b_val = int(data.split('/')[1])
    copy = [x for x in programs]
    temp = copy[a_val]
    copy[a_val] = copy[b_val]
    copy[b_val] = temp
    return copy

def swap_p(programs, data):
    a_val = data.split('/')[0]
    b_val = data.split('/')[1]
    return swap_i(programs, '{}/{}'.format(programs.index(a_val), programs.index(b_val)))

functions = {
        's': spin,
        'x': swap_i,
        'p': swap_p,
        }

with open(sys.argv[1], 'r') as f: dance_moves = f.read().replace('\n', '').split(',')

def execute(programs, dance):
    return functions[dance[0]](programs, dance[1:])

ll = LinkedList(programs)
original_ll = LinkedList(programs)
for dance in dance_moves:
    ll.execute(dance)

count = 1
while original_ll.join() != ll.join():
    for dance in dance_moves:
        ll.execute(dance)
    count += 1


for i in range(0, 1000000000 % count):
    for dance in dance_moves:
        ll.execute(dance)
print(ll.join())
