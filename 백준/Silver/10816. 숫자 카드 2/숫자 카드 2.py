from collections import defaultdict

N = int(input())
src_cards = list(map(int, input().split()))
card_cnt = defaultdict(int)
set_cards = []

for card in src_cards:
    if str(card) not in card_cnt.keys():
        set_cards.append(card)
    card_cnt[str(card)] += 1

M = int(input())
tar_cards = list(map(int, input().split()))

set_cards.sort()


def binary_search(tar, set_cards, card_cnt):
    left = 0
    right = len(set_cards) - 1

    while left <= right:
        mid = (left + right) // 2
        if set_cards[mid] == tar:
            return card_cnt[str(tar)]
        elif set_cards[mid] > tar:
            right = mid - 1
        else:
            left = mid + 1

    return 0


for tar in tar_cards:
    print(binary_search(tar, set_cards, card_cnt), end=" ")
