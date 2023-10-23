from collections import deque
import sys
input = sys.stdin.readline
n = int(input())
matrix = []

dx = [-1, 0, 1, 0]
dy = [0, -1, 0, 1]

dist = [[-1 for i in range(n)] for j in range(n)]

shark_size = 2
fish_cnt = 0
time_cnt = 0
fish = []  # 거리, 크기, dx, dy
sx = 0
sy = 0
for i in range(n):
    arr = list(map(int, input().split()))
    matrix.append(arr)
    for j in range(n):
        if arr[j] == 9:
            sx = i
            sy = j


def update(tar):
    global matrix, sx, sy, shark_size, fish_cnt, time_cnt
    fx = sx + tar[2]
    fy = sy + tar[3]
    fish_dist = tar[0]
    fish_cnt += 1
    if shark_size == fish_cnt:
        shark_size += 1
        fish_cnt = 0
    time_cnt += fish_dist
    matrix[sx][sy] = 0
    sx = fx
    sy = fy

    matrix[sx][sy] = 9



def bfs(cx, cy):
    global dist, fish
    dist[cx][cy] = 0
    queue = deque([(cx, cy)])
    while queue:
        cx, cy = queue.popleft()
        for i in range(4):
            nx = cx + dx[i]
            ny = cy + dy[i]
            if nx < 0 or ny < 0 or nx >= n or ny >= n or matrix[nx][ny] > shark_size:
                continue
            if dist[nx][ny] == -1:
                dist[nx][ny] = 0
                dist[nx][ny] += dist[cx][cy] + 1
                queue.append((nx, ny))
                if 0 < matrix[nx][ny] < shark_size:
                    fish.append((dist[nx][ny], matrix[nx][ny], nx - sx, ny - sy))
while True:
    fish = []
    dist = [[-1 for i in range(n)] for j in range(n)]
    bfs(sx, sy)

    fish.sort(key=lambda x: (x[0], x[2], x[3]))
    if len(fish) == 0:
        break

    update(fish[0])
    fish.clear()
    dist.clear()
    # for i in range(n):
    #     for j in range(n):
    #         print(matrix[i][j],end=" ")
    #     print("")
    # print("--------------")
print(time_cnt)