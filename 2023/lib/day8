score_key = {"2": 1,"3": 2,"4": 3, "5": 4,"6": 5,"7": 6, "8": 7,"9": 8,"T": 9,
             "J": 10, "Q": 11, "K": 12, "A": 13}

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
       
def cat_hands(grouped_cards, original_cards):
    hands = {"five": [], "four": [], "FH": [], "three": [], "two": [], "one": [], "high": []}
    for i, grouped_card in enumerate(grouped_cards):
        grp_len = len(grouped_card)
        grp_max = max([len(x) for x in grouped_card])        
        if grp_len == 1:
            hands["five"].append(original_cards[i])
        elif grp_len == 2 and grp_max == 4:
            hands["four"].append(original_cards[i])
        elif grp_len == 2 and grp_max == 3:
            hands["FH"].append(original_cards[i])      
        elif grp_len == 3 and grp_max == 3:
            hands["three"].append(original_cards[i])      
        elif grp_len == 3 and grp_max == 2:
            hands["two"].append(original_cards[i])        
        elif grp_len == 4 and grp_max == 2:
            hands["one"].append(original_cards[i])        
        elif grp_len == 5:
            hands["high"].append(original_cards[i])
    return hands
   
def score_hand(hand):
    return (score_key[hand[4]] +
            score_key[hand[3]] * 10 ** 2 +
            score_key[hand[2]] * 10 ** 4 +
            score_key[hand[1]] * 10 ** 6 +
            score_key[hand[0]] * 10 ** 8 )          

def score_hands(hands):
    grouped_cards = [group_cards(x[0]) for x in hands]
    categorized_hands = cat_hands(grouped_cards, hands)
    for key in categorized_hands.keys():
        categorized_hands[key].sort(key=lambda x: x[2])
    sum_score = 0
    card_num = 1
    for key in ["high", "one", "two", "three", "FH", "four", "five"]:
        for card in categorized_hands[key]:
            sum_score += card_num * int(card[1])
            card_num += 1
    return sum_score 

if __name__ == "__main__":

    with open('./prompt.txt') as f:
        lines = parse_lines(f)
    
    hands = [(x.split()[0], x.split()[1], score_hand(x.split()[0])) for x in lines]
    print(score_hands(hands))
