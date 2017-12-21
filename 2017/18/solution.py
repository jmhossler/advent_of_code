import sys

from threading import Thread, Lock

mutex = Lock()

ono = False

def is_waiting():
    return ono

def am_waiting():
    global ono
    ono = True

def am_done():
    global ono
    ono = False
cancel_val = False
def should_cancel():
    return cancel_val

def cancel():
    global cancel_val
    cancel_val = True

with open(sys.argv[1], 'r') as f: data = f.read().split('\n')[:-1]

def assembler(our_queue, their_queue, id):
    registers = {'p': id, 'count': 0}

    def snd(args):
        their_queue.append(get_val(args[0]))
        registers['count'] += 1
        return 1

    def set(args):
        registers[args[0]] = get_val(args[1])
        return 1

    def add(args):
        registers[args[0]] = get_val(args[0]) + get_val(args[1])
        return 1

    def mul(args):
        registers[args[0]] = get_val(args[0]) * get_val(args[1])
        return 1

    def mod(args):
        registers[args[0]] = get_val(args[0]) % get_val(args[1])
        return 1

    def rcv(args):
        if len(their_queue) == 0 and len(our_queue) == 0:
            cancel()
            return 1

        while len(our_queue) == 0 and not should_cancel():
            pass
        if should_cancel():
            return 1
        registers[args[0]] = our_queue.pop(0)
        return 1

    def jgz(args):
        if get_val(args[0]) > 0:
            return get_val(args[1])
        return 1

    def get_val(value):
        if value in registers:
            return registers[value]
        elif value.isalpha():
            registers[value] = 0
            return 0
        return int(value)

    instructions = {
            'snd': snd,
            'set': set,
            'add': add,
            'mul': mul,
            'mod': mod,
            'rcv': rcv,
            'jgz': jgz,
            }

    index = 0
    while index >= 0 and index < len(data) and not should_cancel():
        command = data[index].split(' ')
        print(command)
        instruction = instructions[command[0]]
        index += instruction(command[1:])
    print('{}: {}'.format(id, registers['count']))

queue_a = []
queue_b = []

a = Thread(target=assembler, args=(queue_a, queue_b, 0))
b = Thread(target=assembler, args=(queue_b, queue_a, 1))
threads = [a, b]
a.start()
b.start()

a.join()
b.join()

