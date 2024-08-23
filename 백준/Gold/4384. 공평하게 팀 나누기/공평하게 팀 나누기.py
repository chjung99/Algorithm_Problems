N = int(input())
weights = []
for _ in range(N):
    weights.append(int(input()))
total_weight = sum(weights)

dp = [[0 for _ in range(total_weight+1)]for __ in range(N//2+1)]
dp[0][0] = 1

for i in range(N):
    for j in range(N//2, -1, -1):
        for k in range(total_weight, weights[i]-1, -1):
            dp[j][k] = dp[j][k] | dp[j-1][k-weights[i]]

min_diff = -1
ans = []
for i in range(total_weight, -1, -1):
    if dp[N//2][i]:
        another = total_weight - i
        diff = abs(i-another)
        if min_diff == -1 or min_diff > diff:
            min_diff = diff
            ans = [i, another]
ans.sort()
print(ans[0], ans[1])




