with open('2017/Day18/d18input.txt') as i:
    commands = []
    for line in i:
        commands.append(line.strip())


def parse_command(command):
    func = command[:3]
    targets = command[4:].split(' ')
    return func, targets

def sound(x, registers, sounds):
    try:
        x = int(x)
    except:
        x = registers[x]
    return sounds + [x]

def two_value(cmd, x, y, registers):
    try:
        y = int(y)
    except:
        y = registers[y]
    if cmd == 'set':
        registers[x] = y
    elif cmd == 'add':
        registers[x] += y
    elif cmd == 'mul':
        registers[x] *= y
    elif cmd == 'mod':
        registers[x] %= y
    else:
        print("Uknown command {}".format(cmd, x, y))

def recover(x, registers, sounds):
    try:
        x = int(x)
    except:
        x = registers[x]
    if x != 0:
        return sounds[-1]
    return False

def jump(x, y, registers):
    try:
        x = int(x)
    except:
        x = registers[x]
    if x <= 0:
        return 1
    try:
        return int(y)
    except:
        return registers[y]

def initialized_registers(commands):
    registers = {}
    for command in commands:
        for reg in command.split(' '):
            if len(reg) == 1 and reg.isalpha():
                registers[reg] = 0
    return registers

def main(commands):
    sounds = []
    cmd_index = 0
    registers = initialized_registers(commands)
    recovered = False
    while not recovered:
        func, targets = parse_command(commands[cmd_index])
        if func == 'jgz':
            cmd_index += jump(*targets, registers)
            continue
        elif len(targets) > 1:
            two_value(func, *targets, registers)
        elif func == 'snd':
            sounds = sound(*targets, registers, sounds)
        elif func == 'rcv':
            recovered = recover(*targets, registers, sounds)
        else:
            print("Unrecognized func {}".format(func, targets))
        cmd_index += 1
    return recovered

print(main(commands))
    
def yield_main(commands, inp):
    sounds = []
    cmd_index = 0
    registers = initialized_registers(commands)
    recovered = False
    while not recovered:
        func, targets = parse_command(commands[cmd_index])
        if func == 'jgz':
            cmd_index += jump(*targets, registers)
            continue
        elif len(targets) > 1:
            two_value(func, *targets, registers)
        elif func == 'snd':
            yield sound(*targets, registers, sounds)
        elif func == 'rcv':
            registers[targets[0]] = int(inp)
        else:
            print("Unrecognized func {}".format(func, targets))
        cmd_index += 1
    return recovered