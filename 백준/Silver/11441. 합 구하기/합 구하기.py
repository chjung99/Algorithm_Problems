N = int(input())
arr = list(map(int ,input().split()))
M = int(input())
dp = [0 for _ in range(N+1)]

for i in range(1, N+1):
    dp[i] += arr[i-1] + dp[i-1]

for _ in range(M):
    s, e = map(int, input().split())
    print(dp[e]-dp[s-1])