with open('2017/Day5/d5input.txt') as i:
    instructions = []
    for line in i:
        instructions.append(int(line))

def part1(instructions):
    current = 0
    steps = 0
    while True:
        instructions[current] += 1
        current += instructions[current] -1
        steps += 1
        if current < 0 or current >= len(instructions):
            return steps


print(part1([0,3,0,1,-3]))

#print(part1(instructions))

def part2(instructions):
    current = 0
    steps = 0
    while True:
        if instructions[current] > 2:
            instructions[current] -= 1
            current += instructions[current] + 1
        else:
            instructions[current] += 1
            current += instructions[current] - 1
        steps += 1
        if current < 0 or current >= len(instructions):
            return steps


print(part2([0,3,0,1,-3]))

print(part2(instructions))