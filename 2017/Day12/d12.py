from copy import deepcopy

with open('2017/Day12/d12input.txt') as i:
    groups = {}
    for line in i:
        left, right = line.strip().split(' <-> ')
        right = list(right.split(', '))
        groups[left] = right

def part1(groups):
    visited = set()
    to_visit = deepcopy(groups['0'])
    while len(to_visit) > 0:
        visitor = to_visit.pop()
        if visitor in visited:
            continue
        to_visit += groups[visitor]
        visited.add(visitor)
    return len(visited)

print(part1(groups))

def part2(groups, total):
    if len(groups) == 0:
        return total
    visited = set()
    to_visit = groups[list(groups.keys())[0]]
    while len(to_visit) > 0:
        visitor = to_visit.pop()
        if visitor in visited:
            continue
        to_visit += groups[visitor]
        visited.add(visitor)
    for v in visited:
        del(groups[v])
    return part2(groups, total + 1)

print(part2(groups,0))