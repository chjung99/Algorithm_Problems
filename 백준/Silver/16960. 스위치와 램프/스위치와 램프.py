N, M = map(int, input().split())
visit = [0 for _ in range(M+1)]
arr = []
flag = 0
for i in range(N):
    tmp = list(map(int, input().split()))
    arr.append(tmp[1:])
    for e in arr[i]:
        visit[e] += 1


def turnoff(idx, arr, visit):
    result = True
    for e in arr[idx]:
        visit[e] -= 1
        if visit[e] <= 0:
            result = False
    for e in arr[idx]:
        visit[e] += 1
    return result


for i in range(N):
    if turnoff(i, arr, visit):
        flag = 1
        break
print(flag)