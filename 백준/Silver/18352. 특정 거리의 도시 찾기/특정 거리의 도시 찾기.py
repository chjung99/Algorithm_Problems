from collections import deque

N, M, K, X = map(int, input().split())
adj = [[] for _ in range(N + 1)]
dist = [-1 for _ in range(N + 1)]
for _ in range(M):
    A, B = map(int, input().split())
    adj[A].append(B)


def bfs(X, K, adj, dist):
    ret = []
    dist[X] = 0
    queue = deque([X])
    while queue:
        cur = queue.popleft()
        for nxt in adj[cur]:
            if dist[nxt] != -1:
                continue
            queue.append(nxt)
            dist[nxt] = dist[cur] + 1
            if dist[nxt] == K:
                ret.append(nxt)
    return ret


answer = bfs(X, K, adj, dist)
if len(answer) == 0:
    print("-1")
else:
    answer.sort()
    for num in answer:
        print(num)
