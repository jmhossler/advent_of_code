from itertools import combinations


def get_checksum(box_ids):
    return count_matching(2, box_ids) * count_matching(3, box_ids)


def count_matching(n, box_ids):
    return len([box_id for box_id in box_ids if contains_matching(n, box_id)])


def contains_matching(n, box_id):
    letters = {c: box_id.count(c) for c in set(c for c in box_id)}
    return any(True for letter in letters if letters[letter] == n)


def get_common_letters(box_ids):
    return next(remove_difference(a, b)
                for a, b in combinations(box_ids, 2)
                if count_differences(a, b) == 1)


def count_differences(a, b):
    return sum(1 for i, j in zip(a, b) if i != j)


def remove_difference(a, b):
    return ''.join(i for i, j in zip(a, b) if i == j)


if __name__ == '__main__':
    with open('input') as input_file:
        BOX_IDS = input_file.readlines()
    print(get_checksum(BOX_IDS))
    print(get_common_letters(BOX_IDS))
