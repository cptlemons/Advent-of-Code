from functools import reduce

p_inp = (list(range(256)),[197,97,204,108,1,29,5,71,0,50,2,255,248,78,254,63])

p_test = (list(range(5)),[3,4,1,5])

def part1(c_list, lengths):
    cur_indx = 0
    skip = 0
    for length in lengths:
        if cur_indx + length > len(c_list):
            selection = c_list[cur_indx:] + c_list[:length - len(c_list[cur_indx:])]
            selection = selection[::-1]
            back = selection[:len(c_list) - cur_indx]
            front = selection[len(c_list) - cur_indx:]
            c_list = front + c_list[len(front):-len(back)] + back
        else:
            selection = c_list[cur_indx:cur_indx + length]
            selection = selection[::-1]
            c_list = c_list[:cur_indx] + selection + c_list[cur_indx + length:]
        cur_indx += length + skip
        if cur_indx >= len(c_list):
            cur_indx -= len(c_list)
        skip += 1
    return c_list        


def part2(c_list, lengths, cur_indx, skip):
    for length in lengths:
        if cur_indx + length > len(c_list):
            selection = c_list[cur_indx:] + c_list[:length - len(c_list[cur_indx:])]
            selection = selection[::-1]
            back = selection[:len(c_list) - cur_indx]
            front = selection[len(c_list) - cur_indx:]
            c_list = front + c_list[len(front):-len(back)] + back
        else:
            selection = c_list[cur_indx:cur_indx + length]
            selection = selection[::-1]
            c_list = c_list[:cur_indx] + selection + c_list[cur_indx + length:]
        cur_indx += length + skip
        while cur_indx >= len(c_list):
            cur_indx -= len(c_list)
        skip += 1
    return c_list, cur_indx, skip  

def part2_length(str_length):
    suffix = [17, 31, 73, 47, 23]
    ascii_length = []
    for char in str_length:
        ascii_length.append(ord(char))
    return ascii_length + suffix


def dense_hash(nums):
    dense = []
    for chunk in [nums[i:i+16] for i in range(0, len(nums), 16)]:
        dense.append(reduce(lambda i, j: int(i) ^ int(j), chunk))
    return dense

def hex_repr(nums):
    hexes = []
    for num in nums:
        conv = hex(num)[2:]
        if len(conv) < 2:
            conv = '0' + conv
        hexes.append(conv)
    return hexes

def knot_hash(seed):
    skip = cur_indx = 0
    c_list = list(range(256))

    ascii_length = part2_length(seed)
    for i in range(64):
        c_list, cur_indx, skip = part2(c_list, ascii_length, cur_indx, skip)
        print(i)

    dense = dense_hash(c_list)
    hexes = hex_repr(dense)

    return ''.join(hexes)

