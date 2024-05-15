from collections import deque
import sys

input = sys.stdin.readline

N, M = map(int, input().split())
bridges = [[] for _ in range(N + 1)]

for _ in range(M):
    A, B, C = map(int, input().split())
    bridges[A].append([B, C])
    bridges[B].append([A, C])

S, E = map(int, input().split())


def bfs(C):
    visit = [0 for _ in range(N+1)]
    visit[S] = 1
    queue = deque([S])
    while queue:
        cur = queue.popleft()
        for nxt, nc in bridges[cur]:
            if visit[nxt] or nc < C:
                continue
            visit[nxt] = 1
            queue.append(nxt)
    return visit[E]


def binary_search(S, E, bridges):
    left = 1
    right = 10 ** 9
    answer = 0
    while left <= right:
        mid = (left + right) // 2
        if bfs(mid):
            answer = max(answer, mid)
            left = mid + 1
        else:
            right = mid - 1
    return answer


print(binary_search(S, E, bridges))