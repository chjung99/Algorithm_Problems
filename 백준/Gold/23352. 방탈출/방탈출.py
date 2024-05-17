from collections import deque

N, M = map(int, input().split())
dist = [[-1 for _ in range(M)]for __ in range(N)]
board = []
password = -1
max_v = -1
dx = [0, 1, 0, -1]
dy = [1, 0, -1, 0]
for _ in range(N):
    board.append(list(map(int, input().split())))


def init_dist(dist):
    for i in range(N):
        for j in range(M):
            dist[i][j] = -1


def out_of_range(x, y):
    return x < 0 or x >= N or y < 0 or y >= M


def bfs(sx, sy, dist):
    local_max = 0
    dist[sx][sy] = 0
    queue = deque([[sx, sy]])
    while queue:
        cx, cy = queue.popleft()
        for i in range(4):
            nx, ny = cx + dx[i], cy + dy[i]
            if out_of_range(nx, ny) or dist[nx][ny] != -1 or board[nx][ny] == 0:
                continue
            dist[nx][ny] = dist[cx][cy] + 1
            if local_max < dist[nx][ny]:
                local_max = dist[nx][ny]
            queue.append([nx, ny])
    return local_max

def get_end(local_max, dist):
    ex, ey = 0, 0
    max_idx = -1
    for i in range(N):
        for j in range(M):
            if dist[i][j] == local_max:
                if board[i][j] > max_idx:
                    max_idx = board[i][j]
                    ex, ey = i, j
    return ex, ey


for i in range(N):
    for j in range(M):
        if board[i][j] == 0:
            continue
        init_dist(dist)
        sx, sy = i, j
        local_max = bfs(sx, sy, dist)

        ex, ey = get_end(local_max, dist)
        local_pw = board[sx][sy] + board[ex][ey]

        if local_max > max_v:
            max_v = local_max
            password = local_pw
            # password = max(password, local_pw)
        elif local_max == max_v:
            password = max(password, local_pw)
if password == -1:
    print(0)
else:
    print(password)