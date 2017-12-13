with open('2017/Day1/d1input.txt') as i:
    inp_nums = [int(n) for n in list((i.readline()))]
def part1(nums):
    total = 0

    for i in range(len(nums)-1):
        if nums[i] == nums[i+1]:
            total += nums[i]

    if nums[0] == nums[-1]:
        total += nums[0]

    return total
def part2(nums):
    total = 0
    halfway = len(nums) // 2

    for i in range(len(nums)-1):
        if i < halfway and nums[i] == nums[halfway + i]:
            total += nums[i]*2

    return total


test_cases1 = (
    (1122,3),
    (1111,4),
    (1234,0),
    (91212129,9)
)

for case in test_cases1:
    t_nums = [int(n) for n in list(str(case[0]))]
    if part1(t_nums) != case[1]:
        print("Failed case: {} with {}".format(case,part1(t_nums)))
    else:
        print("Passed")

test_cases2 = (
    (1212,6),
    (1221,0),
    (123425,4),
    (123123,12),
    (12131415,4)
)

for case in test_cases2:
    t_nums = [int(n) for n in list(str(case[0]))]
    if part2(t_nums) != case[1]:
        print("Failed case: {} with {}".format(case,part2(t_nums)))
    else:
        print("Passed")

print(part1(inp_nums))

print(part2(inp_nums))