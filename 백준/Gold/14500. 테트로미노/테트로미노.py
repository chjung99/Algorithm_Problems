N, M = map(int, input().split())
board = []
visit = [[0 for _ in range(M)] for __ in range(N)]
dx = [0, 1, 0, -1]
dy = [1, 0, -1, 0]
for _ in range(N):
    board.append(list(map(int, input().split())))

answer = 0


def out_of_range(x, y):
    return x < 0 or x >= N or y < 0 or y >= M


def dfs(cx, cy, visit, depth, local_sum):
    global answer
    if depth == 4:
        answer = max(answer, local_sum)
        return
    for i in range(4):
        nx, ny = cx + dx[i], cy + dy[i]
        if out_of_range(nx, ny) or visit[nx][ny]:
            continue
        visit[nx][ny] = 1
        dfs(nx, ny, visit, depth + 1, local_sum + board[nx][ny])
        visit[nx][ny] = 0


for i in range(N):
    for j in range(M):
        visit[i][j] = 1
        dfs(i, j, visit, 1, board[i][j])
        visit[i][j] = 0

        for p in range(4):
            cnt = 1
            local_sum = board[i][j]

            for q in range(4):
                nx, ny = i + dx[q], j + dy[q]
                if p == q or out_of_range(nx, ny):
                    continue
                cnt += 1
                local_sum += board[nx][ny]

            if cnt == 4:
                answer = max(answer, local_sum)

print(answer)
