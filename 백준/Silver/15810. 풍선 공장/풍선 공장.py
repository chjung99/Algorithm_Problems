N, M = map(int, input().split())
arr = list(map(int, input().split()))

left = 1
right = 1000000 * 1000000


def func(t):
    cnt = 0
    for x in arr:
        cnt += t // x
    return cnt >= M


while left + 1 < right:
    mid = (left + right) // 2
    if func(mid):
        right = mid
    else:
        left = mid
print(right)
