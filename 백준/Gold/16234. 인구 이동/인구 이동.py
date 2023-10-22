from collections import deque
import sys



dx = [0, 1, 0, -1]
dy = [1, 0, -1, 0]

input = sys.stdin.readline


n, L, R = map(int, input().split())
visit = [[0 for i in range(n)] for j in range(n)]
matrix = []


def bfs(cx, cy):
    global visit
    queue = deque([(cx, cy)])
    ret = list([(cx, cy)])
    visit[cx][cy] = 1
    while queue:
        cx, cy = queue.popleft()
        for i in range(4):
            nx = cx + dx[i]
            ny = cy + dy[i]
            if nx < 0 or ny < 0 or nx >= n or ny >= n:
                continue

            if visit[nx][ny] == 0:
                d = matrix[nx][ny] - matrix[cx][cy]
                if d < 0:
                    d *= -1
                if L <= d <= R:
                    visit[nx][ny] = 1
                    queue.append((nx, ny))
                    ret.append((nx, ny))

    return ret

def check(cx, cy):
    global visit

    queue = deque([(cx, cy)])
    visit[cx][cy] = 1
    while queue:
        cx, cy = queue.popleft()
        for i in range(4):
            nx = cx + dx[i]
            ny = cy + dy[i]
            if nx < 0 or ny < 0 or nx >= n or ny >= n:
                continue
            if visit[nx][ny] == 0:
                d = matrix[nx][ny] - matrix[cx][cy]
                if d < 0:
                    d *= -1
                if L <= d <= R:
                    queue.append((nx, ny))
                    return True
    return False


for i in range(n):
    matrix.append(list(map(int, input().split())))

cnt = 0

while True:
    flag = False
    visit = [[0 for i in range(n)] for j in range(n)]
    new_matrix = matrix.copy()
    for i in range(n):
        for j in range(n):
            if not visit[i][j]:
                union = bfs(i, j)
                length = len(union)
                if length > 1:
                    flag = True
                s = 0
                for x, y in union:
                    s += matrix[x][y]

                for x, y in union:
                    new_matrix[x][y] = s // length
    if not flag:

        break
    matrix = new_matrix.copy()
    cnt += 1
print(cnt)
