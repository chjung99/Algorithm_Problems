import heapq

N, M, K, X = map(int, input().split())
graph = [[] for _ in range(N + 1)]
dist = [float("inf") for _ in range(N + 1)]

for i in range(M):
    A, B = map(int, input().split())
    graph[A].append((1, B))


def dijstra(start):
    q = []
    heapq.heappush(q, (0, start))
    dist[start] = 0
    while q:
        cur_dist, cur = heapq.heappop(q)

        if cur_dist > dist[cur]:
            continue

        for nxt_dist, nxt in graph[cur]:
            if dist[nxt] > cur_dist + nxt_dist:
                dist[nxt] = cur_dist + nxt_dist
                heapq.heappush(q, (dist[nxt], nxt))


dijstra(X)

if dist.count(K) == 0:
    print("-1")
else:
    for i in range(1, N+1):
        if dist[i] == K:
            print(i)
