from collections import deque

N, M, x, y, K = map(int, input().split())

A = []

for i in range(N):
    A.append(list(map(int, input().split())))

orders = list(map(int, input().split()))

horizon = deque([0, 0, 0])
vertical = deque([0, 0, 0, 0])

dx = [0, 0, 0, -1, 1]
dy = [0, 1, -1, 0, 0]

cx, cy = x, y
for order in orders:
    nx, ny = cx + dx[order], cy + dy[order]
    if nx >= N or nx < 0 or ny >= M or ny < 0:
        continue

    if order == 1:
        horizon.rotate(1)
        htmp = horizon.popleft()
        horizon.appendleft(vertical.pop())
        vertical.append(htmp)
        vertical[1] = horizon[1]
    elif order == 2:
        horizon.rotate(-1)
        htmp = horizon.pop()
        horizon.append(vertical.pop())
        vertical.append(htmp)
        vertical[1] = horizon[1]
    elif order == 3:
        vertical.rotate(-1)
        horizon[1] = vertical[1]
    else:
        vertical.rotate(1)
        horizon[1] = vertical[1]

    if A[nx][ny] == 0:
        A[nx][ny] = vertical[-1]
    else:
        vertical[-1] = A[nx][ny]
        A[nx][ny] = 0
    cx, cy = nx, ny
    print(horizon[1])
