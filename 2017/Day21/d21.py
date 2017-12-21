
puzzle_start = [['.','#','.'],['.','.','#'],['#','#','#']]

def part1(puzzle):
    for _ in range(5):
        if len(puzzle) % 2 == 0:
            puzzle = div2(puzzle)
        else:
            puzzle = div3(puzzle)

def div2(puzzle):
    sections = []
    for col in range(len(puzzle) % 2):
        for row in range(len(puzzle) % 2):
            first_row = puzzle[col][row:row+2]
            second_row = puzzle[col+1][row:row+2]
            sections.append([first_row, second_row])
    return sections

def div3(puzzle):
    sections = []
    for col in range(len(puzzle) // 3):
        for row in range(len(puzzle) // 3):
            first_row = puzzle[col][row:row+3]
            second_row = puzzle[col+1][row:row+3]
            third_row = puzzle[col+2][row:row+3]
            sections += [first_row, second_row, third_row]
    return sections

    
    
    