def break_down_polymer(code):
    current_check = 0
    while current_check < len(code) - 1:
        if can_be_removed(code, current_check):
            code = code[:current_check] + code[current_check+2:]
            if current_check != 0:
                current_check -= 1
        else:
            current_check += 1
    return code

def can_be_removed(code, current_index):
    a, b = code[current_index], code[current_index+1]
    return (a.lower() == b.lower()
            and ((a.islower() and b.isupper())
                 or (a.isupper() and b.islower())))

def improve_polymer(code):
    letters = {x.lower() for x in code}
    smallest_code = code
    for letter in letters:
        temp_code = strip(code, letter)
        reacted = break_down_polymer(temp_code)
        if len(smallest_code) > len(reacted):
            smallest_code = reacted
    return smallest_code

def strip(code, letter):
    return ''.join(x for x in code if x.lower() != letter)

def solution(input_file):
    with open(input_file) as f:
        code = f.read()[:-1]
    return break_down_polymer(code), improve_polymer(code)
