inp_num = 312051

def part1(inp_num):
    distance = 0
    current = 1
    ring = 1
    while True:
        if inp_num <= current:
            break
        current += 8*ring
        ring += 1
    pos = list(reversed(range(ring))) + list(range(1,ring-1))
    indx = (current - inp_num) % len(pos)
    return ring+pos[indx]-1

test_cases1 = (
    (1,0),
    (12,3),
    (23,2),
    (26,5),
    (1024,31)
)


for case in test_cases1:
    ans = part1(case[0])
    if ans == case[1]:
        print("Passed")
    else:
        print("Failed case {} with {}".format(case,ans))

print(part1(inp_num))

def part2(inp_num):
    #done in excel
    pass