N, M = map(int, input().split())
H = list(map(int, input().split()))
acc_sum = [0 for _ in range(N)] # 0~N-1
tot_sum = [0 for _ in range(N)]
for _ in range(M):
    a, b, k = map(int, input().split())
    acc_sum[a-1] += k
    if b < N:
        acc_sum[b] -= k

tot_sum[0] = acc_sum[0]
for i in range(1, N):
    tot_sum[i] = tot_sum[i-1] + acc_sum[i]
for i in range(N):
    print(H[i]+tot_sum[i], end=" ")