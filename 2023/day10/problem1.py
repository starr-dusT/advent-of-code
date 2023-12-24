flow = {
    "|": [(0,1), (0,-1)],
    "-": [(-1,0), (1,0)],
    "L": [(0,1), (1,0)],
    "J": [(0,1), (-1,0)],
    "7": [(0,-1), (-1,0)],
    "F": [(0, -1), (1,0)]
}

# stop criminal scum! >:(
def legal_moves(curr_pos, curr_char):
    return [add_pos(curr_pos, x) for x in flow[curr_char]]

def add_pos(pos1, pos2):
    return (pos1[0]+pos2[0], pos1[1]+pos2[1])

def process_move(rows, curr_pos, last_pos):
    curr_char = rows[curr_pos[1]][curr_pos[0]]
    if curr_char == '.':
        return 0
    elif curr_char == 'S':
        return 1
    else:
        # legal moves, but don't go backwards
        posb_pos = [x for x in legal_moves(curr_pos, curr_char) if x != last_pos]
        if len(posb_pos) == 1:
            # next legal position
            return (posb_pos[0][0], posb_pos[0][1])
        else:
            return 0

def process_start(rows, start_pos):
    pipe = []
    start_char = rows[start_pos[1]][start_pos[0]]
    if start_char == '.' or start_char == 'S':
        return 0
    for last_pos in legal_moves(start_pos, start_char):
        int_pipe = []
        curr_pos = start_pos
        while curr_pos != 0 and curr_pos != 1: 
            next_pos = process_move(rows, curr_pos, last_pos)
            int_pipe.append(rows[curr_pos[1]][curr_pos[0]])
            last_pos = curr_pos
            curr_pos = next_pos
        if curr_pos:
            if len(pipe):
                pipe.extend(int_pipe[1:])
            else:
                int_pipe.reverse()
                pipe.extend(int_pipe)
    if len(pipe):
        return int((len(pipe)-1)/2)
    else:
        return 0

if __name__ == "__main__":
    # read lines boi
    with open("./prompt.txt") as f:
        rows = f.read().strip().split('\n')
    # make 0 at "bottom"
    rows.reverse()

    # loop over all characters
    for i, row in enumerate(rows):
        for j, char in enumerate(row):
            # determine if we found a continous pipe and return half length
            ans = process_start(rows, (j, i))
            # we found it!
            if ans:
                break
        # we found it!
        if ans:
            break
    print(ans)
