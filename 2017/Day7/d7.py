import sys
sys.setrecursionlimit(10000)

def part1():
    with open('2017/Day7/d7input.txt') as i:
        bots = []
        above = []
        for line in i:
            try:
                left, right = line.strip().split(' -> ')
            except ValueError:
                continue
            bots.append(left.split(' ')[0])
            for r in right.split(', '):
                above.append(r)

    for b in bots:
        if b in above:
            continue
        return b

print(part1())

def part2():
    with open('2017/Day7/d7input.txt') as i:
        bot_weights = {}
        above = {}
        for line in i:
            try:
                left, right = line.strip().split(' -> ')
                bot, weight = left.strip().split(' ')
                bot_weights[bot] = int(weight.strip('(').strip(')'))
                above[bot] = right.split(', ')
            except ValueError:
                bot, weight = line.strip().split(' ')
                bot_weights[bot] = int(weight.strip('(').strip(')'))

    bot_weights['rfkvap'] -= 9
    print(bot_weights['rfkvap'])

    #kept replacing until I found one that balanced then knew the bot below it was the issue
    for bot in above['vtpoo']:
        print(bot)
        print(get_weight(bot,bot_weights,above))

def get_weight(bot,bot_weights,above):
    weight = bot_weights[bot]
    if bot not in above:
        return weight
    for ab in above[bot]:
        weight += get_weight(ab,bot_weights,above)
    return weight

print(part2())