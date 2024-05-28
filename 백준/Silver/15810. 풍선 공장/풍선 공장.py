N, M = map(int, input().split())
A = list(map(int, input().split()))
left = 0
right = min(A) * M


def check(t):
    cnt = 0
    for x in A:
        cnt += t // x
    return cnt >= M


while left < right:
    mid = (left + right) // 2
    if check(mid):
        right = mid
    else:
        left = mid + 1
print(left)
# print(left, right, check(right))


# 3 3
# 1 1 1
# 0 0 False

# 3 8
# 5 7 3
# 14 14 True