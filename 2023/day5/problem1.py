def parse_lines(f):
    lines = []
    for line in f.readlines():
        lines.append(line)
    return lines

def get_seeds(s):
    return [int(x) for x in s.split(": ")[1].split(" ")]

   

def parse_almac(almac):
    almac_dict = dict()
    index = [i for i, v in enumerate(almac) if ":" in v]
    for a, z in list(zip(index[:], index[1:])) + [(index[-1], len(almac))]:
        almac_id = almac[a].split(" ")[0]
        almac_dict[almac_id] = dict()
        for i in range(a+1, z):
            line_list = [int(x) for x in almac[i].split(" ")]
            values = [x for x in range (line_list[0], line_list[0] + line_list[2])]
            keys = [x for x in range (line_list[1], line_list[1] + line_list[2])]
            for key, value in zip(keys, values):
                almac_dict[almac_id][key] = value
    return almac_dict

if __name__ == "__main__":

    with open('input/prompt.txt') as f:
        lines = parse_lines(f)

    lines = [x.replace('\n', '') for x in lines if x != '\n']
    seeds = get_seeds(lines[0])
    parsed_almac = parse_almac(lines[1:])
    small_loc = -1
    for seed in seeds:
        tmp_key = seed
        for key in parsed_almac.keys():
            try: 
                tmp_key = parsed_almac[key][tmp_key] 
            except:
                pass
        if small_loc < 0:
            small_loc = tmp_key
        elif small_loc > tmp_key:
            small_loc = tmp_key
    print(small_loc)
