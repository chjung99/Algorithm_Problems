N, K = map(int, input().split())
arr = list(map(int, input().split()))
ans = 1
L, R = 0, 0
cnt = [0 for _ in range(10 ** 5 + 1)]
cnt[arr[R]] += 1
cur_len = 1
while True:
    if R == N - 1 and R == L:
        break
    if R == N - 1:
        if cnt[arr[R]] <= K:
            ans = max(ans, cur_len)
        L += 1
        cur_len -= 1
    elif cnt[arr[R]] <= K:
        ans = max(ans, cur_len)
        R += 1
        cnt[arr[R]] += 1
        cur_len += 1

    else:
        cnt[arr[L]] -= 1
        L += 1
        cur_len -= 1

print(ans)

