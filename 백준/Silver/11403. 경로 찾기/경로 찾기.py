from collections import deque

N = int(input())
adj_matrix = []
ans_matrix = [[0 for _ in range(N)]for __ in range(N)]

for _ in range(N):
    adj_matrix.append(list(map(int, input().split())))


def bfs(src, adj_matrix, ans_matrix):
    visit = [0 for _ in range(N)]
    visit[src] = 1
    queue = deque([src])
    while queue:
        cur = queue.popleft()
        for i in range(N):
            if adj_matrix[cur][i] and not visit[i]:
                visit[i] = 1
                ans_matrix[src][i] = 1
                queue.append(i)
            if adj_matrix[cur][i] and visit[i] and cur != src and i == src:
                ans_matrix[src][src] = 1


for i in range(N):
    bfs(i, adj_matrix, ans_matrix)

for i in range(N):
    for j in range(N):
        print(ans_matrix[i][j], end=" ")
    print()

