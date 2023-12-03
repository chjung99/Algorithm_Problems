import math

n = int(input())
arr = list(map(int, input().split()))
b, c = map(int, input().split())

ans = n

for a in arr:
    if a > b:
        # print(math.ceil(((a-b) / c)))
        ans += math.ceil(((a-b) / c))
print(ans)


