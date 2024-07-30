from itertools import combinations

N, M = map(int, input().split())
guitars = []

for _ in range(N):
    x, y = input().split()
    guitars.append(y)

# 10 2^10 50
answer = -1
num = 0
for i in range(1, N+1):
    cases = combinations(guitars, i)
    for case in cases:
        visit = [0 for _ in range(M)]
        for play in case:
            for j in range(M):
                if play[j] == 'Y':
                    visit[j] = 1
        if num < sum(visit):
            answer = i
            num = sum(visit)
print(answer)
