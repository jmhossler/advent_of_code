import sys
from common import read_strings


TIME_MAP = {chr(x+65): x+1 for x in range(0, 26)}

def get_step_order(restrictions):
    dependency_graph = parse_dependencies(restrictions)
    root = get_no_dependents(dependency_graph)
    queue = sorted(root)
    step_order = []
    while queue:
        current_place = queue.pop(0)
        step_order.append(current_place)
        for x in get_dependents(dependency_graph, current_place):
            if requirements_satisfied(dependency_graph, step_order, x):
                queue.append(x)
        queue.sort()
    return ''.join(step_order)

def get_step_time(restrictions, workers, offset):
    dependency_graph = parse_dependencies(restrictions)
    root = get_no_dependents(dependency_graph)
    queue = sorted(root)
    total_seconds = 0
    worker_role = {x: '.' for x in range(0, workers)}
    completed = []
    for i in range(0, workers):
        if worker_role[i] == '.' and queue:
            current_place = queue.pop(0)
            worker_role[i] = [current_place, TIME_MAP[current_place] + offset]
        elif not queue and worker_role[i] != '.':
            if worker_role[i][1] != 0:
                worker_role[i][1] -= 1
            elif worker_role[i][1] == 0:
                current_place = worker_role[i][0]
                completed.append(current_place)
                worker_role[i] = '.'
                for x in get_to_be_queued(dependency_graph, current_place, completed):
                    queue.append(x)
    while workers_working(worker_role):
        print('{} {} {}'.format(
            total_seconds,
            ' '.join('{}'.format(worker_role[i]) for i in range(0, workers)),
            completed))
        for i in range(0, workers):
            if worker_role[i] != '.':
                if worker_role[i][1] != 1:
                    worker_role[i][1] -= 1
                else:
                    current_place = worker_role[i][0]
                    completed.append(current_place)
                    worker_role[i] = '.'
                    for x in get_to_be_queued(dependency_graph, current_place, completed):
                        queue.append(x)
                    queue.sort()
        for j in range(0, workers):
            if queue and worker_role[j] == '.':
                current_place = queue.pop(0)
                worker_role[j] = [current_place, TIME_MAP[current_place] + offset]
        total_seconds += 1
        queue.sort()
    return total_seconds

def workers_working(worker_role):
    return [x for x in worker_role if worker_role[x] != '.']

def get_to_be_queued(graph, current_place, completed):
    return [x
            for x in get_dependents(graph, current_place)
            if requirements_satisfied(graph, completed, x)]

def parse_dependencies(restrictions):
    dependencies = {}
    for restriction in restrictions:
        split_restriction = restriction.split(' ')
        a, b = split_restriction[1], split_restriction[7]
        if b not in dependencies:
            dependencies[b] = []
        if a not in dependencies:
            dependencies[a] = []
        dependencies[b].append(a)
    return dependencies

def get_dependents(graph, x):
    return [node for node in graph if x in graph[node]]

def get_no_dependents(graph):
    return [x for x in graph if not graph[x]]

def requirements_satisfied(graph, completed, node):
    for x in graph[node]:
        if x not in completed:
            return False
    return True


def solution(file_name):
    return get_step_order(read_strings(file_name)), get_step_time(read_strings(file_name), 5, 60)

if __name__ == '__main__':
    print(solution(sys.argv[1]))
