n = int(input())
arr = []
dp = [0 for _ in range(n)]

for _ in range(n):
    k = int(input())
    arr.append(k)

if n >= 3:
    dp[0] = arr[0]
    dp[1] = arr[0] + arr[1]
    dp[2] = max(arr[0] + arr[1], max(arr[0] + arr[2], arr[1] + arr[2]))
elif n == 2:
    dp[0] = arr[0]
    dp[1] = arr[0] + arr[1]
else:
    dp[0] = arr[0]

for i in range(3, n):
    dp[i] = max(dp[i - 1], max(dp[i - 3] + arr[i - 1] + arr[i], dp[i - 2] + arr[i]))
print(dp[n - 1])
