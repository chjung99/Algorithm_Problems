# 입력
N, H, W = map(int, input().split())
# 문자열 입력
arr = ['?' for _ in range(N*W)]
for _ in range(H):
    stmp = list(input())
    for i in range(N*W):
        if stmp[i] != '?':
            arr[i] = stmp[i]
for i in range(0, N*W, W):
    ctmp = '?'
    for j in range(i, i+W):
        if arr[j] != '?':
            ctmp = arr[j]
    print(ctmp, end="")