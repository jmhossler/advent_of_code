def get_data_from_line(line):
    data = line.split(' ')
    return {
            'name': data[0],
            'weight': int(data[1].replace('\n', '')[1:-1]),
            'stack': get_stacked(data),
            }

def get_stacked(data):
    if len(data) < 2:
        return []
    return [x.replace('\n','').replace(',','') for x in data[3:len(data)]]

def get_weights(nodes, node):
    if len(node['stack']) == 0:
        return {}
    weights = {}
    for node_on_node in node['stack']:
        sub_weights = get_weights(nodes, nodes[node_on_node])
        weight = nodes[node_on_node]['weight'] + sum([sub_weights[val] for val in sub_weights])
        weights[node_on_node] = weight
    return weights

def get_bad_weight(nodes, node):
    if len(node['stack']) == 0:
        return {}
    weights = get_weights(nodes, node)
    if equal([weights[weight] for weight in weights]):
        return node['name'], node['weight']
    bad_weight = get_unequal([weights[weight] for weight in weights])
    bad_node = get_name_of_weight(weights, bad_weight)
    print(weights)
    return get_bad_weight(nodes, nodes[bad_node])

def get_name_of_weight(weights, bad_weight):
    for weight in weights:
        if weights[weight] == bad_weight:
            return weight

def get_unequal(weights):
    for x in weights:
        copy = [weight for weight in weights]
        copy.remove(x)
        if x not in copy:
            return x
    return 0

def equal(weights):
    for x in weights:
        print(weights)
        copy = [weight for weight in weights]
        copy.remove(x)
        if x not in copy:
            return False
    return True


nodes = {}
with open('input', 'r') as f:
    for line in f:
        node = get_data_from_line(line)
        nodes[node['name']] = node

nodes_with_stuff = [nodes[node] for node in nodes if len(nodes[node]['stack']) > 0]
nodes_on_nodes = []
for node in nodes_with_stuff:
    for thing in node['stack']:
        if thing not in nodes_on_nodes:
            nodes_on_nodes.append(thing)

base_node = [node for node in nodes_with_stuff if node['name'] not in nodes_on_nodes][0]

print(get_bad_weight(nodes, base_node))
