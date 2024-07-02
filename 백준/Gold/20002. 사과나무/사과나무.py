N = int(input())
arr = []
dp = [[0 for _ in range(N)]for __ in range(N)]
for _ in range(N):
    arr.append(list(map(int, input().split())))
answer = -1001


# def func(x, y, z):
#     sum = 0
#     for i in range(x, x+z+1):
#         for j in range(y, y+z+1):
#             sum += arr[i][j]
#     return sum
def func(x, y, z):
    sum = 0
    if x != 0 and y != 0:
        sum = dp[x+z][y+z] - (dp[x+z][y-1]+dp[x-1][y+z])+dp[x-1][y-1]
    elif x == 0 and y!=0:
        sum = dp[x+z][y+z] - (dp[x+z][y-1])
    elif y == 0 and x!=0:
        sum = dp[x+z][y+z] - (dp[x-1][y+z])
    else:
        sum = dp[x+z][y+z]
    return sum


def out_of_range(x, y):
    return x >= N or y >= N

for i in range(N):
    for j in range(N):
        if i == 0 and j == 0:
            dp[i][j] = arr[i][j]
        elif i == 0:
            dp[i][j] = arr[i][j] + dp[i][j-1]
        elif j == 0:
            dp[i][j] = arr[i][j] + dp[i - 1][j]
        else:
            dp[i][j] = arr[i][j] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
# for i in range(N):
#     for j in range(N):
#         print(dp[i][j], end=" ")
#     print()

for i in range(N):
    for j in range(N):
        for k in range(N):
            if out_of_range(i+k, j+k):
                continue
            answer = max(func(i,j,k), answer)

print(answer)
