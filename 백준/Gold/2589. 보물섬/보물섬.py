from collections import deque

N, M = map(int, input().split())
dist = [[0 for _ in range(M)] for __ in range(N)]
arr = []
global_max = -1
dx = [0, 1, 0, -1]
dy = [1, 0, -1, 0]

for _ in range(N):
    arr.append(list(input()))
def init_dist(dist):
    for i in range(N):
        for j in range(M):
            dist[i][j] = -1


def out_of_range(x, y):
    return x >= N or x < 0 or y >= M or y < 0


def bfs(cx, cy, dist):
    local_max = 0
    dist[cx][cy] = 0
    queue = deque([[cx, cy]])
    while queue:
        cx, cy = queue.popleft()
        for i in range(4):
            nx, ny = cx + dx[i], cy + dy[i]
            if out_of_range(nx, ny) or dist[nx][ny] != -1 or arr[nx][ny] == 'W':
                continue
            dist[nx][ny] = dist[cx][cy] + 1
            queue.append([nx, ny])

            if local_max < dist[nx][ny]:
                local_max = dist[nx][ny]

    return local_max


for i in range(N):
    for j in range(M):
        if arr[i][j] == 'W':
            continue
        init_dist(dist)
        global_max = max(global_max, bfs(i, j, dist))
print(global_max)
