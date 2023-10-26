# 순열 nPr
import sys


# def input():
#     return sys.stdin.readline().rstrip()


n, m = map(int, input().split())
choose = [0 for _ in range(10)]
used = [0 for _ in range(10)]

def dfs(cnt):
    if cnt == m:
        for idx in range(cnt):
            print(choose[idx], end=' ')
        print()
    for i in range(1, n+1):
        if not used[i]:
            used[i] = 1
            choose[cnt] = i
            dfs(cnt+1)
            choose[cnt] = 0
            used[i] = 0
dfs(0)