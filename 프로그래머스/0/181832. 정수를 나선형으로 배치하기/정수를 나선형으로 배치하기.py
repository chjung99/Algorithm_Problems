def out_of_range(x, y, n):
    return x < 0 or x >= n or \
        y < 0 or y >= n

def solution(n):
    arr = [[0 for _ in range(n)]for __ in range(n)]
    visit = [[0 for _ in range(n)]for __ in range(n)]
    dx = [0, 1, 0, -1]
    dy = [1, 0, -1, 0]
    cx, cy, cd = 0, 0, 0
    cnt = 0
    T = n ** 2

    while True:
        visit[cx][cy] = 1
        cnt += 1
        arr[cx][cy] = cnt
        
        if cnt == T:
            break

        nx, ny = cx + dx[cd], cy + dy[cd]
        
        if out_of_range(nx, ny, n) or visit[nx][ny]:
            cd = (cd + 1) % 4
            nx, ny = cx + dx[cd], cy + dy[cd]

        cx, cy = nx, ny

    return arr