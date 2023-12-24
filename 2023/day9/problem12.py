# notes:
# recursion?

def find_next_diff(seq):
    int_seq = [j-i for i, j in zip(seq[:-1], seq[1:])]
    if int_seq == len(int_seq) * [int_seq[0]]:
        return seq[-1] + int_seq[0]
    else:
        return seq[-1] + find_next_diff(int_seq)

if __name__ == "__main__":
    # read lines boi
    with open("./prompt.txt") as f:
        lines = f.read().strip().split('\n')

    for part in ["Part 1", "Part 2"]:
        next_sum = 0
        for line in lines:
            seq = [int(n) for n in line.split()]
            if part == "Part 2": seq.reverse()
            next_sum += find_next_diff(seq)
        print(part + ": ", next_sum)
