with open('2017/Day8/d8input.txt') as i:
    instructions = []
    for line in i:
        instructions.append(line.strip())

def part1(instructions):
    registers = {}
    for instruc in instructions:
        command, condition = instruc.split(' if ')
        reg1, cmd, value1 = command.split(' ')
        value1 = int(value1)
        if reg1 not in registers:
            registers[reg1] = 0
        if cmd == 'dec':
            value1 = -value1

        reg2, cond, value2 = condition.split(' ')
        value2 = int(value2)
        if reg2 not in registers:
            registers[reg2] = 0
        if cond == '>':
            if registers[reg2] <= value2:
                continue
        elif cond == '<':
            if registers[reg2] >= value2:
                continue
        elif cond == '>=':
            if registers[reg2] < value2:
                continue
        elif cond == '<=':
            if registers[reg2] > value2:
                continue
        elif cond == '==':
            if registers[reg2] != value2:
                continue
        elif cond == '!=':
            if registers[reg2] == value2:
                continue
        else:
            print('Unknown condition {}', cond)
            return False
        registers[reg1] += value1
    return registers


test = ['b inc 5 if a > 1','a inc 1 if b < 5','c dec -10 if a >= 1','c inc -20 if c == 10']

registers = part1(instructions)
print(max(registers.values()))

def part2(instructions):
    registers = {}
    highest = 0
    for instruc in instructions:
        command, condition = instruc.split(' if ')
        reg1, cmd, value1 = command.split(' ')
        value1 = int(value1)
        if reg1 not in registers:
            registers[reg1] = 0
        if cmd == 'dec':
            value1 = -value1

        reg2, cond, value2 = condition.split(' ')
        value2 = int(value2)
        if reg2 not in registers:
            registers[reg2] = 0
        if cond == '>':
            if registers[reg2] <= value2:
                continue
        elif cond == '<':
            if registers[reg2] >= value2:
                continue
        elif cond == '>=':
            if registers[reg2] < value2:
                continue
        elif cond == '<=':
            if registers[reg2] > value2:
                continue
        elif cond == '==':
            if registers[reg2] != value2:
                continue
        elif cond == '!=':
            if registers[reg2] == value2:
                continue
        else:
            print('Unknown condition {}', cond)
            return False
        registers[reg1] += value1
        highest = max(highest,max(registers.values()))
    return highest

print(part2(instructions))