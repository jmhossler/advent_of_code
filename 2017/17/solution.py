import sys

with open(sys.argv[1], 'r') as f: steps = int(f.read())

buf = [0]
curr_index = 0

class Node:
    def __init__(self, val):
        self.val = val
        self.next = None

    def set_next(self, node):
        self.next = node

    def get_val(self):
        return self.val

class ll:
    def __init__(self, val):
        self.head = Node(val)
        self.head.set_next(self.head)
        self.zero = self.head
        self.length = 1

    def insert(self, val, index):
        new_node = Node(val)

        curr = 0
        probe = self.head
        index = index % self.length
        while curr < index:
            probe = probe.next
            curr += 1
        next_node = probe.next
        probe.set_next(new_node)
        new_node.set_next(next_node)

        self.length += 1
        self.head = new_node

    def get_next(self):
        return self.head.next.get_val()

buff = ll(0)

for i in range(1, 50000001):
    buff.insert(i, steps)

print(buff.get_next())
print(buff.zero.next.get_val())
