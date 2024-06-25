N, M, K = map(int, input().split())
arr = [[i + M * j for i in range(1, M + 1)] for j in range(N)]
dp = [[0 for _ in range(M)] for __ in range(N)]
dp2 = [[0 for _ in range(M)] for __ in range(N)]
dp[0][0] = 1
if K == 0:
    for i in range(N):
        for j in range(M):
            if i>0:
                dp[i][j] += dp[i-1][j]
            if j>0:
                dp[i][j] += dp[i][j-1]
    print(dp[N-1][M-1])
else:
    px = (K-1) // M
    py = (K-1) % M
    # print(px, py)
    for i in range(px+1):
        for j in range(py+1):
            if i>0:
                dp[i][j] += dp[i-1][j]
            if j>0:
                dp[i][j] += dp[i][j-1]
    ans = dp[px][py]

    dp2[px][py]= 1
    for i in range(px, N):
        for j in range(py, M):
            if i>0:
                dp2[i][j] += dp2[i-1][j]
            if j>0:
                dp2[i][j] += dp2[i][j-1]
    ans *= dp2[N-1][M-1]
    print(ans)