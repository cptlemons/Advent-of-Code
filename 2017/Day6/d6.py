from copy import deepcopy

inp = [10,3,15,10,5,15,5,15,9,2,5,8,5,2,3,6]
test = [0,2,7,0]

def part1(mem_bank):
    mem_bank = deepcopy(mem_bank)
    seen = set()
    while str(mem_bank) not in seen:
        seen.add(str(mem_bank))
        redis = max(mem_bank)
        start = mem_bank.index(redis) + 1
        mem_bank[start-1] = 0
        mem_bank = mem_bank[start:] + mem_bank[:start]
        while redis >= len(mem_bank):
            redis -= len(mem_bank)
            mem_bank = [n+1 for n in mem_bank]
        for i in range(redis):
            mem_bank[i] += 1
        mem_bank = mem_bank[-start:] + mem_bank[:-start]
    return len(seen)

print(part1(test))

print(part1(inp))

def part2(mem_bank):
    mem_bank = deepcopy(mem_bank)
    seen = []
    while str(mem_bank) not in seen:
        seen.append(str(mem_bank))
        redis = max(mem_bank)
        start = mem_bank.index(redis) + 1
        mem_bank[start-1] = 0
        mem_bank = mem_bank[start:] + mem_bank[:start]
        while redis >= len(mem_bank):
            redis -= len(mem_bank)
            mem_bank = [n+1 for n in mem_bank]
        for i in range(redis):
            mem_bank[i] += 1
        mem_bank = mem_bank[-start:] + mem_bank[:-start]
    return len(seen) - seen.index(str(mem_bank))

print(part2(test))

print(part2(inp))

