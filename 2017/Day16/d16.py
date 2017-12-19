from copy import deepcopy

with open('2017/Day16/d16input.txt') as i:
    for line in i:
        moves = line.strip().split(',')

programs = ['a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p']

def check_move(move, programs):
    if move[0] == 's':
        spin = int(move[1:])
        programs = spin_move(spin%16,programs)
    elif move[0] == 'x':
        ex_a, ex_b = move[1:].split('/')
        programs = exchange_move(int(ex_a), int(ex_b), programs)
    elif move[0] == 'p':
        sw_a = move[1]
        sw_b = move[3]     
        programs = swap_move(sw_a, sw_b, programs)
    else:
        print('Something went wrong')
    return programs

def spin_move(spin,programs):
    programs = programs[-spin:] + programs[:-spin]
    return programs

def exchange_move(ex_a, ex_b, programs):
    a = programs[ex_a]
    b = programs[ex_b]
    programs[ex_a] = b
    programs[ex_b] = a
    return programs

def swap_move(sw_a, sw_b, programs):
    a = programs.index(sw_a)
    b = programs.index(sw_b)
    programs[a] = sw_b
    programs[b] = sw_a
    return programs

seen_states = []

d_programs = deepcopy(programs)

for i in range(150):
    for move in moves:
        d_programs = check_move(move, d_programs)
    if str(d_programs) in seen_states:
        repeat_len = len(seen_states) - seen_states.index(str(d_programs))
        offset = seen_states.index(str(d_programs))
        break
    else:
        seen_states.append(str(d_programs))

print(repeat_len, offset)

for i in range(offset):
    for move in moves:
        programs = check_move(move,programs)

for i in range(1000000000%repeat_len):
    for move in moves:
        programs = check_move(move, programs)

print(''.join(programs))