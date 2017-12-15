from functools import reduce

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

    dense = dense_hash(c_list)
    hexes = hex_repr(dense)

    return ''.join(hexes)

t_input = 'stpzcrnm-'


def count_row(inp):
    row = bin(int(inp, 16))[2:]
    return row.count('1')

def part1(seed):
    total = 0
    for i in range(128):
        inp = knot_hash(t_input + str(i))
        total += count_row(inp)
    return total

def make_row(inp):
    row = bin(int(inp, 16))[2:]
    return '{:0>128}'.format(row)

def make_grid(inp):
    grid = ()
    for i in range(128):
        inp = knot_hash(t_input + str(i))
        grid += (list(make_row(inp)),)
    return grid

def find_adj(grid, row, col):
    search_nodes = [(row,col)]
    while len(search_nodes) > 0:
        row, col = search_nodes.pop()
        grid[row][col] = '0'
        if row > 0 and grid[row-1][col] == '1':
            search_nodes += [(row-1, col)]
        if row < 127 and grid[row+1][col] == '1':
            search_nodes += [(row+1, col)]
        if col > 0 and grid[row][col-1] == '1':
            search_nodes += [(row, col-1)]
        if col < 127 and grid[row][col+1] == '1':
            search_nodes += [(row, col+1)]
    return grid

def get_groups(grid):
    groups = 0
    for r in range(128):
        for c in range(128):
            if grid[r][c] == '1':
                find_adj(grid, r, c)
                groups += 1
    return groups

def main(inp):
    grid = make_grid(inp)
    return get_groups(grid)

print(main(t_input))