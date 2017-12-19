maze = []

with open('2017/Day19/d19input.txt') as i:
    for line in i:
        maze.append(list(line.strip()))

cur_pos = (0,maze[0].index('|'))
current_char = '|'


letters = ''
direction = 'd'
steps = 1


while True:
    if direction == 'd':
        for i in range(cur_pos[0] + 1, len(maze)):
            steps += 1
            char = maze[i][cur_pos[1]]
            if char.isalpha():
                letters += char
            if char == '+':
                if maze[i][cur_pos[1] + 1] != '`':
                    direction = 'r'
                    cur_pos = (i, cur_pos[1])
                elif maze[i][cur_pos[1] - 1] != '`':
                    direction = 'l'
                    cur_pos = (i, cur_pos[1])
                break
            if char == '`':
                break
    elif direction == 'u':
        for i in reversed(range(cur_pos[0])):
            steps += 1
            char = maze[i][cur_pos[1]]
            if char.isalpha():
                letters += char
            if char == '+':
                if maze[i][cur_pos[1] + 1] != '`':
                    direction = 'r'
                    cur_pos = (i, cur_pos[1])
                elif maze[i][cur_pos[1] - 1] != '`':
                    direction = 'l'
                    cur_pos = (i, cur_pos[1])
                break
    elif direction == 'r':
        for i in range(cur_pos[1] + 1, len(maze[cur_pos[0]])):
            steps += 1
            char = maze[cur_pos[0]][i]
            if char.isalpha():
                letters += char
            if char == '+':
                if maze[cur_pos[0] - 1][i] != '`':
                    direction = 'u'
                    cur_pos = (cur_pos[0], i)
                elif maze[cur_pos[0] + 1][i] != '`':
                    direction = 'd'
                    cur_pos = (cur_pos[0], i)
                break
    elif direction == 'l':
        for i in reversed(range(cur_pos[1])):
            steps += 1
            char = maze[cur_pos[0]][i]
            if char.isalpha():
                letters += char
            if char == '+':
                if maze[cur_pos[0] - 1][i] != '`':
                    direction = 'u'
                    cur_pos = (cur_pos[0], i)
                elif maze[cur_pos[0] + 1][i] != '`':
                    direction = 'd'
                    cur_pos = (cur_pos[0], i)
                break
    if char == '`':
        break

print(letters, steps - 1)

