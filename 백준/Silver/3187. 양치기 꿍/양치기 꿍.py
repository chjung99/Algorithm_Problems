from collections import deque

R, C = map(int, input().split())
arr = []
visit = [[0 for _ in range(C + 1)] for __ in range(R + 1)]
dx = [1, 0, -1, 0]
dy = [0, 1, 0, -1]
ans_k = 0
ans_v = 0

for _ in range(R):
    arr.append(list(input()))


def bfs(x, y):
    global visit, ans_k, ans_v
    tmp_k = 0
    tmp_v = 0
    visit[x][y] = 1
    queue = deque([[x,y]])
    while queue:
        cur = queue.popleft()
        if arr[cur[0]][cur[1]] == 'v':
            tmp_v += 1
        elif arr[cur[0]][cur[1]] == 'k':
            tmp_k += 1

        for i in range(4):
            nx = cur[0] + dx[i]
            ny = cur[1] + dy[i]
            if nx < 0 or nx >= R or ny < 0 or ny >= C:
                continue
            if arr[nx][ny] != '#' and not visit[nx][ny]:
                queue.append([nx, ny])
                visit[nx][ny] = 1
    if tmp_k > tmp_v:
        ans_k += tmp_k
    else:
        ans_v += tmp_v


for i in range(R):
    for j in range(C):
        if (arr[i][j] == 'v' or arr[i][j] == 'k') and not visit[i][j]:
            bfs(i, j)

print(ans_k, ans_v)
