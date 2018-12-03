import sys

from common import read_strings


def count_overlap(claims):
    filled = populate_map(parse_input(claims))
    return len([x for x in filled if filled[x] > 1])

def populate_map(claims):
    filled = {}
    for claim_id in claims:
        claim = claims[claim_id]
        for x in range(claim[0][0], claim[0][0] + claim[1][0]):
            for y in range(claim[0][1], claim[0][1] + claim[1][1]):
                if (x, y) not in filled:
                    filled[(x, y)] = 0
                filled[(x, y)] += 1
    return filled

def get_0_overlap_id(claims):
    parsed_data = parse_input(claims)
    filled = populate_map(parsed_data)
    return next(claim_id
                for claim_id in parsed_data
                if not has_overlap(filled, parsed_data[claim_id]))

def has_overlap(filled, claim):
    for x in range(claim[0][0], claim[0][0] + claim[1][0]):
        for y in range(claim[0][1], claim[0][1] + claim[1][1]):
            if filled[(x, y)] > 1:
                return True
    return False

def parse_input(claims):
    return {
        _parse_id(claim): [_parse_start_corner(claim), _parse_end_corner(claim)]
        for claim in claims
    }

def _parse_id(claim):
    id_str, _, _, _ = claim.split(' ')
    return int(id_str[1:])

def _parse_start_corner(claim):
    _, _, start_str, _ = claim.split(' ')
    start_x, start_y = start_str[:-1].split(',')
    return _str_to_int_tuple(start_x, start_y)

def _parse_end_corner(claim):
    _, _, _, end_str = claim.split(' ')
    end_x_delta, end_y_delta = end_str.split('x')
    return _str_to_int_tuple(end_x_delta, end_y_delta)

def _str_to_int_tuple(a, b):
    return (int(a), int(b))

def solution(input_file):
    claims = read_strings(input_file)
    return (count_overlap(claims), get_0_overlap_id(claims))

if __name__ == '__main__':
    print(solution(sys.argv[1]))
