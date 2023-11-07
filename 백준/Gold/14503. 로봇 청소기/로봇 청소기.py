from collections import deque
cnt = 0
dx = [-1, 0, 1, 0]
dy = [0, 1, 0, -1]
visit = [[0 for i in range(55)] for j in range(55)]
matrix = []
n, m = map(int, input().split())
x, y, d = map(int, input().split())

for i in range(n):
    matrix.append(list(map(int, input().split())))

queue = deque([[x, y, d]])
visit[x][y] = 1
cnt += 1
while queue:
    # print("===================")
    # for i in range(n):
    #     for j in range(m):
    #         print(visit[i][j], end=" ")
    #     print()
    cx, cy, cd = queue.popleft()

    flag = True
    nd = cd
    for _ in range(4):
        nd = nd-1
        if nd < 0:
            nd = 3
        nx = cx + dx[nd]
        ny = cy + dy[nd]
        if nx < 0 or nx >= n or ny < 0 or ny >= m:
            continue
        if matrix[nx][ny] == 0 and visit[nx][ny] == 0:
            cnt += 1
            visit[nx][ny] = 1
            queue.append([nx, ny, nd])
            flag = False
            break
    if flag:
        rd = (cd+2) % 4
        bx = cx + dx[rd]
        by = cy + dy[rd]

        if bx < 0 or bx >= n or by < 0 or by >= m:
            break
        if matrix[bx][by] == 0:

            queue.append([bx, by, cd])
        else:
            queue.clear()
print(cnt)