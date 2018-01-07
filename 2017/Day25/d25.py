from collections import defaultdict

states = {
    'a0':(1,'r','b'),
    'a1':(1,'l','e'),
    'b0':(1,'r','c'),
    'b1':(1,'r','f'),
    'c0':(1,'l','d'),
    'c1':(0,'r','b'),
    'd0':(1,'r','e'),
    'd1':(0,'l','c'),
    'e0':(1,'l','a'),
    'e1':(0,'r','d'),
    'f0':(1,'r','a'),
    'f1':(1,'r','c')
}

steps = 12459852
num_line = defaultdict(int)
index = 0
state = 'a'

for i in range(steps):
    value = num_line[index]
    state += str(value)
    n_value, direction, state = states[state]
    num_line[index] = n_value
    if direction == 'r':
        index += 1
    else:
        index -= 1
ones = 0

for v in num_line.values():
    if v == 1:
        ones += 1

print(ones)
