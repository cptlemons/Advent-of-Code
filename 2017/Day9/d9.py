
def clean_garbage(string):
    garb_list = list(string)
    clean = []
    while len(garb_list) > 0:
        if garb_list[0] != '<':
            clean += garb_list[0]
            garb_list = garb_list[1:]
        else:
            while garb_list[0] != '>':
                if garb_list[0] != '!':
                    garb_list = garb_list[1:]
                else:
                    garb_list = garb_list[2:]
            garb_list = garb_list[1:]
    for char in clean:
        if char not in ('{','}'):
            clean.remove(char)
    return ''.join(clean)

def score_groups(score_list):
    s = []
    index = 0
    score = 0
    while index < len(score_list):
        token = score_list[index]
        if token == '{':
            s.append(token)
        elif token == '}':
            score += len(s)
            s.pop()
        index += 1
    return score


all_garbage_strings = (
    '<>',
    '<random characters>',
    '<<<<>',
    '<{!>}>',
    '<!!>',
    '<!!!>>',
    '<{o"i!a,<{i<a>'
)

for garb in all_garbage_strings:
    if clean_garbage(garb) != '':
        print('Failed garbage string {}'.format(garb))

test_groups = (
    ('{}',1),
    ('{{{}}}',6),
    ('{{},{}}',5),
    ('{{{},{},{{}}}}',16),
    ('{<a>,<a>,<a>,<a>}',1),
    ('{{<ab>},{<ab>},{<ab>},{<ab>}}',9),
    ('{{<!!>},{<!!>},{<!!>},{<!!>}}',9),
    ('{{<a!>},{<a!>},{<a!>},{<ab>}}',3)
)

for group in test_groups:
    cleaned = clean_garbage(group[0])
    if score_groups(cleaned) != group[1]:
        print("Failed scoring on {}".group)

with open('2017/Day9/d9input.txt') as i:
    string = i.readline().strip()

print(score_groups(clean_garbage(string)))

def clean_garbage_count(string):
    garb_list = list(string)
    chars_cleaned = 0
    while len(garb_list) > 0:
        if garb_list[0] != '<':
            garb_list = garb_list[1:]
        else:
            garb_list = garb_list[1:]
            while garb_list[0] != '>':
                if garb_list[0] != '!':
                    chars_cleaned += 1
                    garb_list = garb_list[1:]
                else:
                    garb_list = garb_list[2:]
            garb_list = garb_list[1:]
    return chars_cleaned

for garb in all_garbage_strings:
    print(clean_garbage_count(garb))

print(clean_garbage_count(string))
