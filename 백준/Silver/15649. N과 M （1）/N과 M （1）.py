# 순열 nPr
ret = []
n, m = map(int, input().split())


def func(v):
    global ret
    if len(v) == m:
        ret.append(v.copy())
        return
    for j in range(1, n+1):
        if j not in v:
            v.append(j)
            func(v)
            v.pop()

for i in range(1, n+1):
    func([i])
for vv in ret:
    for v in vv:
        print(v, end=" ")
    print("")