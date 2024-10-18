def solution(n, money):
    answer = 0
    m = len(money)
    dp = [[0 for _ in range(n+1)]for __ in range(m+1)]
    r =  1000000007
    # dp[i][n] = money의 i번째까지 집합으로 n을 거슬러주는 경우의 수
    
    for i in range(m):
        dp[i][0] = 1
        for x in range(1, n+1):
            if i == 0:
                if x % money[i] == 0:
                    dp[i][x] = 1
            else:
                if x < money[i]:
                    dp[i][x] = dp[i-1][x] 
                else:
                    dp[i][x] = (dp[i-1][x] + dp[i][x-money[i]]) % r
    # for b in dp:
    #     print(b)
    return dp[m-1][n] % r