import itertools

containers = [43, 3, 4, 10, 21, 44, 4, 6, 47, 41, 34, 17, 17, 44, 36, 31, 46, 9, 27, 38]
combinations = [c for i in xrange(1, len(containers)+1) for c in itertools.combinations(containers, i) if sum(p) == 150]
print len(combinations)  # part1
print len([c for c in combinations if len(c) == len(min(combinations, key=lambda x: len(x)))])  # part2
