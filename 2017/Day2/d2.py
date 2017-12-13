with open('2017/Day2/d2input.txt') as i:
    inp_rows = []
    for line in i:
        inp_rows.append([int(n) for n in line.strip().split('\t')])

def part1(rows):
    total = 0
    for row in rows:
        total += max(row) - min(row)
    return total

def part2(rows):
    total = 0
    for row in rows:
        row.sort(reverse=True)
        for i in range(len(row)):
            for j in range(1,len(row)-i):
                if row[i] % row[i+j] == 0:
                    total += row[i] / row[i+j]
    return total

test_cases1 = (
    ([[5,1,9,5],[7,5,3],[2,4,6,8]],18),
)

for case in test_cases1:
    rows = case[0]
    if part1(rows) != case[1]:
        print("Failed case: {} with {}".format(case,part1(rows)))
    else:
        print("Passed")

print(part1(inp_rows))


test_cases2 = (
    ([[5,9,2,8],[9,4,7,3],[3,8,6,5]],9),
)

for case in test_cases2:
    rows = case[0]
    if part2(rows) != case[1]:
        print("Failed case: {} with {}".format(case,part2(rows)))
    else:
        print("Passed")


print(part2(inp_rows))