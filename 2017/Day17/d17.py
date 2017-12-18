p_inp = 335
t_inp = 3

def spinner(inp):
    circ = [0]
    value = 1
    index = 0
    while value < 2018:
        insert = (inp % len(circ)) + index
        if insert >= len(circ):
            insert -= len(circ)
        circ.insert(insert, value)
        index = insert + 1
        value += 1

    return circ[circ.index(2017)+1]

print(spinner(p_inp))

def angry_spinner(inp):
    circ = [0]
    value = 1
    index = 0
    while value <= 50000000:
        insert = (inp % len(circ)) + index
        if insert >= len(circ):
            insert -= len(circ)
        circ.insert(insert, value)
        index = insert + 1
        value += 1
        if value % 100000 == 0:
            print(value)

    return circ[circ.index(0)+1]

from collections import deque

def p2_spinner(inp):
    spinner = deque([0])
    for value in range(1,50000000+1):
        spinner.rotate(-inp)
        spinner.append(value)
    return spinner[spinner.index(0) + 1]

print(p2_spinner(p_inp))

