# 순열 nPr
n, m = map(int, input().split())


def func(v, idx):
    if len(v) == m:
        for vv in v:
            print(vv, end=" ")
        print()
    for i in range(idx, n+1):
        if i not in v:
            v.append(i)
            func(v, i+1)
            v.pop()
for i in range(1, n+1):
    func([i], i+1)