def execute_command(command):
    if command[0] not in registers:
        registers[command[0]] = 0
    if command[1] == 'dec':
        registers[command[0]] -= int(command[2])
    else:
        registers[command[0]] += int(command[2])

commands = []
with open('input', 'r') as f:
    for line in f:
        commands.append(line)

max_val = 0
registers = {}
for command in commands:
    command = command.split(' ')
    executable = command[:3]
    conditional = command[4:]
    str_cond = ' '.join(conditional)
    if conditional[0] not in registers:
        registers[conditional[0]] = 0
    str_cond = str_cond.replace(conditional[0], str(registers[conditional[0]]))
    if (eval(str_cond)):
        execute_command(command)
    max_of_iter = max([registers[x] for x in registers])
    if max_of_iter > max_val:
        max_val = max_of_iter

print(max([registers[x] for x in registers]))
print(max_val)
