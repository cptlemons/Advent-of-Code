test_inp = (65,8921)
p1_inp = (618,814)

def part1(inp):
    a_factor = 16807
    b_factor = 48271
    a_prev = inp[0]
    b_prev = inp[1]
    div = 2147483647
    score = 0
    for i in range(40000000):
        a_num = (a_prev * a_factor) % div
        b_num = (b_prev * b_factor) % div
        a_prev, b_prev = a_num, b_num
        if a_num % 65536 == b_num % 65536:
            score += 1
        if i % 1000000 == 0:
            print("{} million".format(i/1000000))
    return(score)

def part2(inp):
    a_factor = 16807
    b_factor = 48271
    a_prev = inp[0]
    b_prev = inp[1]
    div = 2147483647
    score = 0
    for i in range(5000000):
        a_num = (a_prev * a_factor) % div
        while a_num % 4 != 0:
            a_prev = a_num
            a_num = (a_prev * a_factor) % div
        b_num = (b_prev * b_factor) % div
        while b_num % 8 != 0:
            b_prev = b_num
            b_num = (b_prev * b_factor) % div
        if a_num % 65536 == b_num % 65536:
            score += 1
        a_prev, b_prev = a_num, b_num
        if i % 100000 == 0:
            print(i)
    return(score)

print(part2(p1_inp))