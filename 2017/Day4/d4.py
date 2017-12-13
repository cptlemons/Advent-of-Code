with open('2017/Day4/d4input.txt') as i:
    passphrases = []
    for line in i:
        passphrases.append((line.strip().split(' ')))

def part1(passphrases):
    num_valid = 0
    for phrase in passphrases:
        seen = set()
        valid = True
        for word in phrase:
            if word in seen:
                valid = False
                break
            seen.add(word)
        if valid:
            num_valid += 1
    return num_valid

print(part1(passphrases))

def part2(passphrases):
    num_valid = 0
    for phrase in passphrases:
        seen = set()
        valid = True
        for word in phrase:
            word = ''.join(sorted(word))
            if word in seen:
                valid = False
                break
            seen.add(word)
        if valid:
            num_valid += 1
    return num_valid

print(part2(passphrases))
