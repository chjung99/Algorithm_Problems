# 입력
T = int(input())
for _ in range(T):
    N = int(input())
    dp0 = [0 for __ in range(40+1)]
    dp1 = [0 for __ in range(40+1)]
    dp0[0] = 1
    dp0[1] = 0
    dp1[0] = 0
    dp1[1] = 1
    for i in range(2, N+1):
        dp0[i] = dp0[i-1] + dp0[i-2]
        dp1[i] = dp1[i-1] + dp1[i-2]

    print(f"{dp0[N]} {dp1[N]}")

