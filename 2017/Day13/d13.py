from copy import deepcopy

with open('2017/Day13/d13input.txt') as i:
    layers = []
    for line in i:
        left, right = line.strip().split(': ')
        layers.append((int(left),int(right)))


def part1(layers):
    pico = 0
    layer = 0
    severity = 0 
    while len(layers) > 0:
        if layer == layers[0][0]:
            depth, l_range = layers.pop(0)
            top = 2 * l_range - 2
            if pico % top == 0:
                severity += depth * l_range
        pico += 1
        layer += 1
    return severity

print(part1(deepcopy(layers)))

def part2(layers, pico):
    layer = 0
    while len(layers) > 0:
        if layer == layers[0][0]:
            depth, l_range = layers.pop(0)
            top = 2 * l_range - 2
            if pico % top == 0:
                return False
        pico += 1
        layer += 1
    return True

''' for i in range(0,15000000,2):
    if part2(deepcopy(layers),i):
        print("No detections at {}".format(i))
        break
    if (i) % 10000 == 0:
        print(i) '''


def part2_math(layers):
    for i in range(10000000):
        for layer in layers:
            valid = True
            if (i + layer[0]) % (layer[1]*2 - 2) == 0:
                valid = False
                break
        if valid:
          print("No detections at {}".format(i))
          return None

part2_math(layers)