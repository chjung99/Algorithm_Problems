N = int(input())
arr = list(map(int, input().split()))
M = int(input())
acc_sum = [0 for _ in range(N+1)]
for i in range(1, N+1):
    acc_sum[i] = acc_sum[i-1] + arr[i-1]

for _ in range(M):
    i, j = map(int, input().split())
    print(acc_sum[j]-acc_sum[i-1])