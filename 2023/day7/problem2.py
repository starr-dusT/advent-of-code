# including nerfed J!
score_key = {"2": 2,"3": 3,"4": 4, "5": 5,"6": 6,"7": 7, "8": 8,"9": 9,"T": 10,
             "J": 1, "Q": 12, "K": 13, "A": 14}

def parse_lines(f):
    lines = []
    for line in f.readlines():
        lines.append(line)
    return lines

def group_cards(cards):
    lst_cards = list(cards)
    lst_cards.sort()
    grouped_cards = []
    tmp_group = lst_cards[0]
    for i, card in enumerate(lst_cards[1:]):
        if card == lst_cards[i] or tmp_group == "":
            tmp_group += card
        else:
            grouped_cards.append(tmp_group)
            tmp_group = card
    grouped_cards.append(tmp_group)
    return grouped_cards

def cat_hand(hand):
    grouped_hand = group_cards(hand)
    grp_len = len(grouped_hand)
    grp_max = max([len(x) for x in grouped_hand])        

    if grp_len == 1:                    # five of a kind
        return 6
    elif grp_len == 2 and grp_max == 4: # four of a kind
        return 5
    elif grp_len == 2 and grp_max == 3: # full house
        return 4
    elif grp_len == 3 and grp_max == 3: # three of a kind
        return 3
    elif grp_len == 3 and grp_max == 2: # two pair
        return 2
    elif grp_len == 4 and grp_max == 2: # one pair
        return 1
    elif grp_len == 5:                  # high card
        return 0
    else:
        return 100   

def score_hand(hand):
    return (score_key[hand[4]] +
            score_key[hand[3]] * 10 ** 2 +
            score_key[hand[2]] * 10 ** 4 +
            score_key[hand[1]] * 10 ** 6 +
            score_key[hand[0]] * 10 ** 8 )          

if __name__ == "__main__":

    with open('./prompt.txt') as f:
        lines = parse_lines(f)

    categorized_hands = [[], [], [], [], [], [], []]
     
    for line in lines: 
        hand = (line.split()[0], line.split()[1], score_hand(line.split()[0]))
        best_cat = max([cat_hand(hand[0].replace("J", x)) for x in '23456789TQKA'])
        categorized_hands[best_cat].append(hand)

    for cat in categorized_hands:
        cat.sort(key=lambda x: x[2])

    card_num = 1
    sum_score = 0
    for cat in categorized_hands:
        for hand in cat:
            sum_score += card_num * int(hand[1])
            card_num += 1   
    print(sum_score)
