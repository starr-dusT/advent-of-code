import math

def parse_lines(f):
    lines = []
    for line in f.readlines():
        lines.append(line)
    return lines

def parse_races(lines):
    races = dict()
    count = 0
    for time, dist in zip(lines[0].split()[1:], lines[1].split()[1:]):
        races[count] = [int(time), int(dist)]
        count += 1
    return races

def find_win_times(record_time, record_distance):
    dis = record_time * record_time - 4 * (-1) * (-1 * record_distance)
    sqrt_dis = math.sqrt(abs(dis))
    sol1 = (-1 * record_time + sqrt_dis)/(2*-1) + 1 # Get next whole number after sol
    sol2 = (-1 * record_time - sqrt_dis)/(2*-1) - 1 # Get prev whole number before sol
    return list(range(math.floor(sol1), math.ceil(sol2)+1))

if __name__ == "__main__":

    with open('prompt.txt') as f:
        lines = parse_lines(f)
    races = parse_races(lines)
    product = 1
    for race_num in races.keys():
        product *= len(find_win_times(races[race_num][0], races[race_num][1]))
    print("Part1: ", product)
   
    large_time = int(''.join(lines[0].split()[1:]))
    large_dist = int(''.join(lines[1].split()[1:]))
    win_times = find_win_times(large_time, large_dist)
    print("Part2: ", len(win_times))
