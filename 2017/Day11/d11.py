from copy import deepcopy

with open('2017/Day11/d11input.txt') as i:
    directions = i.readline().strip().split(',')

def part1(directions):
    directions = deepcopy(directions)

    new_directions = remove_opposites(directions)
    count = {}

    for direct in new_directions:
        if direct in count:
            count[direct] += 1
        else:
            count[direct] = 1
    
    return sum(count.values()) - min(count.values()), new_directions

def remove_opposites(directions):
    while 'ne' in directions and 'sw' in directions:
        directions.remove('ne')
        directions.remove('sw')
    while 'nw' in directions and 'se' in directions:
        directions.remove('nw')
        directions.remove('se')
    while 'n' in directions and 's' in directions:
        directions.remove('n')
        directions.remove('s')
    return directions

print(part1(directions)[0])

def part2(directions):
    directions = deepcopy(directions)
    max_distance = 0
    path = []
    while len(directions) > 0:
        step = directions.pop(0)
        path.append(step)
        dist, path = part1(path)
        max_distance = max(max_distance, dist)
    return max_distance

print(part2(directions))