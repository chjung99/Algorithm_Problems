T = int(input())
ans = []
for _ in range(T):
    J, N = map(int, input().split())
    box = []
    cnt = 0
    for __ in range(N):
        r, c = map(int, input().split())
        box.append(r * c)
    box.sort(reverse=True)
    idx = 0
    while J > 0:
        cnt += 1
        J -= box[idx]
        idx += 1
    ans.append(cnt)
for x in ans:
    print(x)
