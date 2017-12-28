from copy import deepcopy
from heapq import heappush, heappop

with open('2017/Day24/d24input.txt') as i:
    connectors = []
    for line in i:
        connectors += [line.strip().split('/')]

def part1(current, bridge, connectors, q):
    bridge = deepcopy(bridge)
    for connector in connectors:
        if current == connector[0]:
            n_connectors = deepcopy(connectors)
            n_connectors.remove(connector)
            heappush(q,(connector[1], bridge + connector, n_connectors))
        elif current == connector[1]:
            n_connectors = deepcopy(connectors)
            n_connectors.remove(connector)
            heappush(q,(connector[0], bridge + connector[::-1], n_connectors))
    return len(bridge), sum(map(int, bridge))


def main(connectors):
    q = []
    seen = set() #probably not actually optimizing but making it worse
    length = 0
    p1_strength = 0
    p2_strength = 0
    heappush(q,('0', [], connectors))
    while len(q) > 0:
        current, bridge, connectors = heappop(q)
        if (current, ''.join(bridge), ''.join(str(c) for c in connectors)) in seen:
            continue
        seen.add((current, ''.join(bridge), ''.join(str(c) for c in connectors)))
        n_length, n_strength = part1(current, bridge, connectors, q)
        p1_strength = max(p1_strength,n_strength)
        if n_length > length:
            p2_strength = n_strength
            length = n_length
        elif n_length == length:
            p2_strength = max(p2_strength, n_strength)
    print(p1_strength, p2_strength)


main(connectors) #expect ~5min runtime