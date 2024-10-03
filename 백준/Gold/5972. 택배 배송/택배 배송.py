import heapq


def dijkstra(start, dist, adj_list):
    dist[start] = 0
    pq = []
    heapq.heappush(pq, [0, start])
    while pq:
        cur_weight, cur = heapq.heappop(pq)
        if dist[cur] != cur_weight:
            continue
        for nxt, weight in adj_list[cur]:
            if dist[nxt] > dist[cur] + weight:
                dist[nxt] = dist[cur] + weight
                heapq.heappush(pq, [dist[nxt], nxt])
    return


N, M = map(int, input().split())
adj_list = [[] for _ in range(N + 1)]
dist = [float("inf") for _ in range(N + 1)]

for _ in range(M):
    ai, bi, ci = map(int, input().split())
    adj_list[ai].append([bi, ci])
    adj_list[bi].append([ai, ci])

dijkstra(1, dist, adj_list)
print(dist[N])
