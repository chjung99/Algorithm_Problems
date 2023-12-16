import math
N = int(input())


def mix(k, cards):
    picked = cards
    size = len(cards)
    remain = []
    for i in range(1, k+2):
        idx = 2 ** (k - i + 1)
        new_pick = picked[-idx:]
        if idx + 1 == size:
            no_picked = [picked[0]]
        else:
            no_picked = picked[:-idx]

        cards = new_pick + no_picked + remain

        remain = no_picked + remain
        picked = new_pick

    return cards


tar = list(map(int, input().split()))
M = int(math.log2(N))
for i in range(1, M+1):
    flag = False
    for j in range(1, M+1):
        k_1 = i
        k_2 = j

        src = [i for i in range(1, N + 1)]
        mixed_cards = mix(k_1, src)
        mixed_cards = mix(k_2, mixed_cards)

        if mixed_cards == tar:
            print(k_1, k_2)
            flag = True
            break
    if flag:
        break
