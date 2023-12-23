def build_map(lines):
    nodes = {}
    for line in lines:
        split = line.split(" ")
        nodes[split[0]] = (split[2][1:-1], split[3][:-1])
    return nodes

if __name__ == "__main__":

    # read lines boi
    with open("./prompt.txt") as f:
        lines = f.read().strip().split('\n')

    turns = lines[0]
    nodes = build_map(lines[2:])

    k = 'AAA'

    count = 0
    while 1:
        turn = turns[count % len(turns)]
        count += 1
        if turn == "L":
            k = nodes[k][0]
        else:
            k = nodes[k][1]
        if k == "ZZZ":
            print(count)
            break
