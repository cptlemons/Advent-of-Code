with open('2017/Day23/d23input.txt') as i:
    commands = []
    for line in i:
        commands += [line.strip().split(' ')]


registers = {}

for char in 'abcdefgh':
    registers[char] = 0

def part1(commands, registers):
    indx = 0
    mul_count = 0
    while indx >= 0 and indx < len(commands):
        cmd, first, second = commands[indx]
        try:
            second = int(second)
        except ValueError:
            second = registers[second]

        if cmd == 'set':
            registers[first] = second
        elif cmd == 'sub':
            registers[first] -= second
        elif cmd == 'mul':
            registers[first] *= second
            mul_count += 1
        elif cmd == 'jnz':
            try:
                first = int(first)
            except ValueError:
                first = registers[first]
            if first != 0:
                indx += second
                continue
        indx += 1

def part2():
    with open('2017/Day23/primes.txt') as i:
        primes = i.readline().strip().split(',')
        print(len(primes), primes)
    total = 1001
    for p in primes:
        if (int(p) - 107900) % 17 == 0:
            total -= 1
    print(total)
part2()

'''
exit condition line 29 g == 0
line 28 sub c from g so c == g
line 27 set g to b, so b == c == g
line 26 adds 1 to h (answer) as long as f != 0 on 25
line 24 g != 0
line 23 g == b
line 22 set g to d, so g == b == d
line 21 d += 1
line 20 g == 0
line 19 g == b
line 18 g == b == e
line 17 e += 1
line 16 set f = 0 if g == 0 on 15
line 14 g -= b
line 13 g*= e
line 12 g = d

first loop changes at e = 53950
12 - set g to d (d = 2) - 2
13 - mul g by e - 107900
14 - sub b from g - 0
15 - no jump
16 - f = 0
17 - e += 1  - 53951
18 - g = e - 53951
19 - sub b from g - not 0
At this points it continues until e = 107900 after line 17
18 - g = e = 107900
19 - sub b from g - 0
20 - no jump
21 - d += 1 (3)
22 - set g to d (3)
23 - sub b from g - not 0
24 - jump back 18 into the loop
At this point it continues until d = 107900 after line 21
22 - set g to d (107900)
23 - sub b from g - 0
24 - no jump
25 - no jump (f set to 0 above)
26 - h += 1
27 - set g to b (107900)
28 - sub c from b (107900 - 124900) -17000
29 - jump to 31
31 - add 17 to b 107900 + 17 = 107917
32 - jump back into loop at very beginning preserving b and c only

It takes 1000 additional loops (1001 total) for b == c with a chance each loop to h += 1

Conditions for h being incremented by 1 (f == 0):
    b == g * e and g == d
    e is incrementing by 1 starting at 2 until e == b
    after e == b, d is incrementing by 1 starting at 2 until d == b
    only then does h get incremented by 1 if f == 0
    
    so b % d == 0
    therefore h = 1001 - # of primes between 107900,124900

'''