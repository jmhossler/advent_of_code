def break_down_polymer(code):
    polymer_chain = []
    for letter in code:
        if polymer_chain:
            a = polymer_chain.pop()
            if not can_be_removed(a, letter):
                polymer_chain.append(a)
                polymer_chain.append(letter)
        else:
            polymer_chain.append(letter)

    return ''.join(polymer_chain)

def can_be_removed(a, b):
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
