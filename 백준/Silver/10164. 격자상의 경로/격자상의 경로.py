N, M, K = map(int, input().split())
arr = [[i + M * j for i in range(1, M + 1)] for j in range(N)]
visit = [[0 for _ in range(M)] for __ in range(N)]
cnt = 0
dx = [1, 0]
dy = [0, 1]


def out_of_range(nx, ny):
    return nx < 0 or nx >= N or ny < 0 or ny >= M


def dfs(is_count, cx, cy, visit, arr):
    global cnt

    if arr[cx][cy] == K:
        is_count = True

    if not is_count and arr[cx][cy] > K:
        return



    if arr[cx][cy] == N * M:
        if is_count:
            cnt += 1
        return
    for i in range(2):
        nx = cx + dx[i]
        ny = cy + dy[i]
        if out_of_range(nx, ny) or visit[nx][ny]:
            continue
        visit[nx][ny] = 1
        dfs(is_count, nx, ny ,visit, arr)
        visit[nx][ny] = 0


visit[0][0] = 1
if K == 0:
    dfs(True,0,0,visit,arr)
else:
    dfs(False, 0, 0 , visit, arr)

print(cnt)


