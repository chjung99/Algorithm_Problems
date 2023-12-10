from collections import deque

N, K = map(int, input().split())
A = deque(list(map(int, input().split())))
robots = deque([False for _ in range(N)])
answer = 0

while A.count(0) < K:
    # 1. 벨트가 각 칸 위에 있는 로봇과 함께 한 칸 회전한다.
    A.rotate(1)
    robots.rotate(1)
    if robots[N - 1]:
        robots[N - 1] = False
    # 2. 로봇이동
    for i in range(N - 2, -1, -1):
        if robots[i] and not robots[i+1] and A[i+1] >= 1:
            robots[i + 1] = robots[i]
            robots[i] = False
            A[i+1] -= 1
    if robots[N - 1]:
        robots[N - 1] = False
    # 3. 로봇 올리기
    if A[0] >= 1 and not robots[0]:
        A[0] -= 1
        robots[0] = True
    answer += 1

    # print(robots)
    # print(A)
print(answer)