from collections import deque

g = [-1]  # idx = 1~4
for i in range(4):
    g.append(deque(map(int, list(input()))))

K = int(input())


def func(n, d, vis):
    global g

    left, right = n - 1, n + 1
    if left >= 1 and vis[left] == 0:
        if g[left][2] != g[n][6]:
            vis[left] = 1
            func(left, d * -1, vis)
            g[left].rotate(d * -1)

    if right <= 4 and vis[right] == 0:
        if g[right][6] != g[n][2]:
            vis[right] = 1
            func(right, d * -1, vis)
            g[right].rotate(d * -1)


for i in range(K):
    n, d = map(int, input().split())
    vis = [0 for _ in range(5)]
    vis[n] = 1
    func(n, d, vis)
    g[n].rotate(d)

score = 0
x = 1
for i in range(1, 5):
    # print(g[i])
    if g[i][0] == 1:
        score += x
    x *= 2
print(score)
