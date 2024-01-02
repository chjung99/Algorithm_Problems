import heapq

n, m, r = map(int, input().split())
items = list(map(int, input().split()))
graph = [[] for _ in range(n + 1)]
answer = 0

for _ in range(r):
    a, b, l = map(int, input().split())
    graph[a].append([l, b])
    graph[b].append([l, a])


def dijkstra(start):
    global answer
    pq = []
    dist = [float("inf") for _ in range(n + 1)]
    dist[start] = 0
    heapq.heappush(pq, [0, start])

    while pq:
        cur_dist, cur = heapq.heappop(pq)
        if cur_dist > dist[cur]:
            continue
        for nxt_dist, nxt in graph[cur]:
            if cur_dist+nxt_dist < dist[nxt]:
                dist[nxt] = cur_dist+nxt_dist
                heapq.heappush(pq, [cur_dist+nxt_dist, nxt])

    tmp = 0
    for i in range(1, n + 1):
        if dist[i] <= m:
            tmp += items[i-1]
    answer = max(answer, tmp)

for i in range(1, n+1):
    dijkstra(i)
print(answer)