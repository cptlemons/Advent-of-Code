from copy import deepcopy
grid = []
extra = ['.'] * 5003
for i in range(2500):
    grid.append(deepcopy(extra))

with open('2017/Day22/d22input.txt') as i:
    for line in i:
        grid.append(list(line.strip()))

def turn(facing, loc):
    up = (-1,0)
    down = (1,0)
    left = (0,-1)
    right = (0,1)
    if loc == 'w':
        return facing
    elif loc == 'f':
        if facing == up:
            return down
        elif facing == down:
            return up
        elif facing == right:
            return left
        return right

    elif facing == up:
        if loc == '.':
            return left
        return right
    elif facing == down:
        if loc == '#':
            return left
        return right
    elif facing == left:
        if loc == '.':
            return down
        return up
    elif facing == right:
        if loc == '#':
            return down
        return up


for i,row in enumerate(grid[2500:]):
    grid[i+2500] = ['.'] * 2500 + row + ['.'] * 2500
for i in range(2500):
    grid.append(deepcopy(extra))

loc = (2512,2512)
facing = (-1,0)
infect = 0

for i in range(10000000):
    facing = turn(facing,grid[loc[0]][loc[1]])
    try:
        char = grid[loc[0]][loc[1]]
        if char == '.':
            grid[loc[0]][loc[1]] = 'w'
        elif char == 'w':
            grid[loc[0]][loc[1]] = '#'
            infect += 1
        elif char == '#':
            grid[loc[0]][loc[1]] = 'f'
        else:
            grid[loc[0]][loc[1]] = '.'
    except:
        print('exception')
        print(loc)
        break

    loc = (loc[0] + facing[0], loc[1] + facing[1])

print(loc,infect)