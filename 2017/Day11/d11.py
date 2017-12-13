with open('2017/Day11/d11input.txt') as i:
    directions = i.readline().strip().split(',')

def part1(directions):
    test_cases = (
        ['ne','ne','ne'],
        ['ne','ne','sw','sw'],
        ['ne','ne','s','s'],
        ['se','sw','se','sw','sw']
    )

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

    for case in test_cases:
        remove_opposites(case)

    new_directions = remove_opposites(directions)
    count = {}

    for direct in new_directions:
        if direct in count:
            count[direct] += 1
        else:
            count[direct] = 1

    print(sum(count.values())-min(count.values()))

part1(directions)