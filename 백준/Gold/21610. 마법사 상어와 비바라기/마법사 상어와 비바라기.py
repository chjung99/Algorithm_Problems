# 좌하단에 비구름 생성
# M번 이동 - 방향 칸 (8방향)
# 모든 구름이 di 방향으로 si 칸 이동
# 비가 내려 구름 있는 칸 물의 양 증가 -> 구름 칸 저장
# 구름 사라짐
# 물이 증가한 칸에 물복사 마법(대각선 방향으로 지금 있는 칸 칸에 거리가 1인 칸에 있는
# 바구니의 수만큼 물이 증가한 칸에 바구니의 물의 양이 증가+)
# 이때 경계를 넘어가는 칸은 대각선 방향으로 거리가 1이 아님
# 바구니에 저장된 물의양이 2이상인 모든 칸에 구름이 생기고, 물의양이 2 줄어듬
# 단 이때 구름이 생기는 칸은 방금 구름이 사라진칸이 아니여야함
from collections import deque

n, m = map(int, input().split())
matrix = []
visit = [[0 for i in range(n)] for j in range(n)]
clouds = deque([[n - 1, 0], [n - 1, 1], [n - 2, 0], [n - 2, 1]])

dx = [0, 0, -1, -1, -1, 0, 1, 1, 1]  # 0/1~8
dy = [0, -1, -1, 0, 1, 1, 1, 0, -1]


def move(d, s):
    global clouds
    tmp = deque([])
    while clouds:
        cx, cy = clouds.popleft()
        for i in range(s):
            cx += dx[d]
            if cx == -1:
                cx = n - 1
            if cx == n:
                cx = 0
        for i in range(s):
            cy += dy[d]
            if cy == -1:
                cy = n - 1
            if cy == n:
                cy = 0
        tmp.append([cx, cy])
    while tmp:
        clouds.append(tmp.popleft().copy())


# 대각선 확인은 2 4 6 8 -> 1 2 3 4 x2

def chk(cx, cy):
    cnt = 0
    for i in range(1, 4 + 1):
        nx = cx + dx[i * 2]
        ny = cy + dy[i * 2]
        if nx < 0 or nx >= n or ny < 0 or ny >= n:
            continue

        if matrix[nx][ny] > 0:
            cnt += 1
    return cnt


for i in range(n):
    matrix.append(list(map(int, input().split())))
for i in range(m):
    d, s = map(int, input().split())
    visit = [[0 for j in range(n)] for k in range(n)]
    move(d, s)
    tmp = deque([])
    while clouds:
        cx, cy = clouds.popleft()
        matrix[cx][cy] += 1
        tmp.append([cx, cy])
        visit[cx][cy] = 1
    while tmp:
        cx, cy = tmp.popleft()
        matrix[cx][cy] += chk(cx, cy)
    for j in range(n):
        for k in range(n):
            if visit[j][k] == 0 and matrix[j][k] >= 2:
                clouds.append([j, k])
                matrix[j][k] -= 2
cnt = 0
for j in range(n):
    for k in range(n):
        cnt += matrix[j][k]
print(cnt)
