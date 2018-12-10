from itertools import cycle


def get_high_score(players, marbles):
    scores = {player: 0 for player in range(1, players+1)}
    player_ids = cycle(player_id for player_id in range(1, players+1))
    play_field = CDLL()
    player = next(player_ids)
    for marble in range(1, marbles+1):
        if marble % 23 == 0:
            scores[player] += marble
            scores[player] += play_field.remove(-7)
        else:
            play_field.add(1, marble)
        player = next(player_ids)
    return max(scores[player] for player in scores)

class CDLL:
    def __init__(self):
        self.head = DLNode(0)
        self.head.next = self.head
        self.head.previous = self.head

    def add(self, offset, value):
        for _ in range(0, offset):
            if offset < 0:
                self.head = self.head.previous
            else:
                self.head = self.head.next
        new_node = DLNode(value)
        new_node.next = self.head.next
        new_node.previous = self.head
        self.head.next.previous = new_node
        self.head.next = new_node
        self.head = new_node

    def remove(self, offset):
        direction = -1 if offset < 0 else 1
        for _ in range(0, offset, direction):
            if offset < 0:
                self.head = self.head.previous
            else:
                self.head = self.head.next
        previous = self.head.previous
        next_v = self.head.next
        value = self.head.value
        previous.next = next_v
        next_v.previous = previous
        self.head = next_v
        return value


    def to_list(self):
        probe = self.head.next
        values = []
        while probe != self.head:
            values.append(probe.value)
            probe = probe.next
        return values

class DLNode:
    def __init__(self, value):
        self.value = value
        self.next = None
        self.previous = None

def solution(filename):
    with open(filename) as f:
        line = f.read().strip().split(' ')
        players, marbles = int(line[0]), int(line[6])
    return get_high_score(players, marbles), get_high_score(players, marbles * 100)
