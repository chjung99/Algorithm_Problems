from collections import deque

n, k = map(int, input().split())
arr = list(map(int, input().split()))


def check(ls):
    cnt = 0
    for a in ls:
        if a == 0:
            cnt += 1
    return cnt < k


last = 2 * n
belt = [i for i in range(n)]  # belt 위 칸의 절대 위치(0 ~ 2n-1)
robots = [0 for _ in range(n)]  # robot이 있는지 상대 위치(0 ~ n-1)
answer = 0
while check(arr):

    # 1. 벨트가 각 칸 위에 있는 로봇과 함께 한 칸 회전

    last = (last - 1) % (2 * n)
    belt.pop()
    belt.insert(0, last)
    # print(robots)
    for i in range(n-1, 0, -1):
        if robots[i-1] == 1:
            robots[i] = robots[i-1]
            robots[i-1] = 0
    # print(robots)
    if robots[n-1] == 1:
        robots[n-1] = 0
    # 2. 로봇 이동
    for i in range(n-2, 0, -1):
        if robots[i] == 0:
            continue

        if arr[belt[i+1]] >= 1 and robots[i+1] == 0:
            robots[i] = 0
            robots[i+1] = 1
            arr[belt[i+1]] -= 1


    # 3. 올리는 위치에 내구도가 0이 아니면 로봇을 올린다
    # print(belt[0], robots[0])
    # print(robots)
    if arr[belt[0]] >= 1 and robots[0] == 0:

        robots[0] = 1
        arr[belt[0]] -= 1
    answer += 1

    # for b in belt:
    #     print(arr[b], end=" ")
    # print("\n")
    # print(belt, arr)
    # print(robots)

print(answer)