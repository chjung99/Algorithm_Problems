import sys

# def input():
#     return sys.stdin.readline().rstrip()

n, m = map(int, input().split())
visit = [[0 for i in range(m+1)] for j in range(n+1)]

dx = [0, -1, -1]
dy = [-1, 0, -1]
cnt = 0


def func(used):
    global visit, cnt

    if used == n * m:
        cnt += 1
        return

    x = used // m + 1
    y = used % m + 1

    func(used + 1)  # x, y에 넴모를 두지 않음
    if visit[x-1][y] == 0 or visit[x][y-1] == 0 or visit[x-1][y-1] == 0:
    # if check(x, y):
        visit[x][y] = 1
        func(used + 1)  # x, y에 넴모를 둠
        visit[x][y] = 0


func(0)
print(cnt)
